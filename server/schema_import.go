package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/encoding"

	"github.com/batchcorp/plumber/pb"
	"github.com/batchcorp/plumber/validate"
)

var (
	ErrEmptyFile = errors.New("file contents cannot be empty")
)

// importGithub imports a github repo or file as a schema
func (s *Server) importGithub(ctx context.Context, importReq *protos.ImportGithubRequest, req *protos.ImportGithubSelectRequest) (*protos.Schema, error) {

	var schema *protos.Schema
	var err error

	switch importReq.Type {
	case protos.SchemaType_SCHEMA_TYPE_PROTOBUF:
		schema, err = s.importGithubProtobuf(ctx, importReq, req)
	case protos.SchemaType_SCHEMA_TYPE_AVRO:
		schema, err = s.importGithubAvro(ctx, importReq, req)
	case protos.SchemaType_SCHEMA_TYPE_JSONSCHEMA:
		schema, err = s.importGithubJSONSchema(ctx, importReq, req)
	default:
		err = validate.ErrInvalidGithubSchemaType
	}

	if err != nil {
		return nil, err
	}

	return schema, nil
}

// importGithubProtobuf is used to import a protobuf schema from a GitHub repository
func (s *Server) importGithubProtobuf(ctx context.Context, importReq *protos.ImportGithubRequest, req *protos.ImportGithubSelectRequest) (*protos.Schema, error) {
	repo, err := parseRepoURL(importReq.GithubUrl)
	if err != nil {
		return nil, errors.Wrap(err, "unable to understand repository URL")
	}

	zipFile, err := s.GithubService.GetRepoArchive(ctx, s.PersistentConfig.GitHubToken, repo.Organization, repo.Name)
	if err != nil {
		return nil, err
	}

	if len(zipFile) == 0 {
		return nil, errors.New("zipFile cannot be empty")
	}

	settings := req.GetProtobufSettings()
	if settings == nil {
		return nil, errors.New("protobuf settings cannot be nil")
	}

	fds, files, err := pb.GetFDFromArchive(zipFile, settings.XProtobufRootDir)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse protobuf zip")
	}

	mdBlob, err := pb.CreateBlob(fds, settings.ProtobufRootMessage)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get message descriptor")
	}
	settings.XMessageDescriptor = mdBlob

	//md, err := pb.GetMDFromDescriptorBlob(mdBlob, req.GetProtobufSettings().ProtobufRootMessage)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to get message descriptor")
	//}

	//inferred, err := jsonschema.InferFromProtobuf("http://plumber.local/schema.json", "Inferred schema", "Inferred from protobuf", md)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to convert Protobuf schema to JSONSchema")
	//}
	var inferred []byte

	inferredJSON, err := json.Marshal(inferred)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal inferred JSONSchema")
	}

	return &protos.Schema{
		Id:   uuid.NewV4().String(),
		Name: importReq.Name,
		Type: protos.SchemaType_SCHEMA_TYPE_PROTOBUF,
		InferredSchema: &encoding.JSONSchemaSettings{
			Schema: inferredJSON,
		},
		Versions: []*protos.SchemaVersion{
			{
				Version: 1,
				Status:  protos.SchemaStatus_SCHEMA_STATUS_ACCEPTED,
				Files:   files,
				Settings: &protos.SchemaVersion_ProtobufSettings{
					ProtobufSettings: settings,
				},
			},
		},
	}, nil
}

// importGithubAvro is used to import an avro schema from a GitHub repository
func (s *Server) importGithubAvro(ctx context.Context, importReq *protos.ImportGithubRequest, req *protos.ImportGithubSelectRequest) (*protos.Schema, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	repo, err := parseRepoURL(importReq.GithubUrl)
	if err != nil {
		return nil, errors.Wrap(err, "unable to understand repository URL")
	}

	schemaData, err := s.GithubService.GetRepoFile(ctx, s.PersistentConfig.GitHubToken, repo.Organization, repo.Name, req.SchemaFileSha)
	if err != nil {
		return nil, err
	}

	//inferred, err := jsonschema.InferFromAvro("http://plumber.local/schema.json", schemaData)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to convert Protobuf schema to JSONSchema")
	//}
	var inferred []byte

	inferredJSON, err := json.Marshal(inferred)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal inferred JSONSchema")
	}

	return &protos.Schema{
		Id:   uuid.NewV4().String(),
		Name: importReq.Name,
		Type: protos.SchemaType_SCHEMA_TYPE_AVRO,
		InferredSchema: &encoding.JSONSchemaSettings{
			Schema: inferredJSON,
		},
		Versions: []*protos.SchemaVersion{
			{
				Version: 1,
				Status:  protos.SchemaStatus_SCHEMA_STATUS_ACCEPTED,
				Files: map[string]string{
					req.SchemaFileName: string(schemaData),
				},
				Settings: &protos.SchemaVersion_AvroSettings{
					AvroSettings: &encoding.AvroSettings{
						Schema: schemaData,
					},
				},
			},
		},
	}, nil
}
func (s *Server) importGithubJSONSchema(ctx context.Context, importReq *protos.ImportGithubRequest, req *protos.ImportGithubSelectRequest) (*protos.Schema, error) {
	if len(importReq.GithubUrl) == 0 {
		return nil, errors.New("Github URL cannot be empty")
	}

	repo, err := parseRepoURL(importReq.GithubUrl)
	if err != nil {
		return nil, errors.Wrap(err, "unable to understand repository URL")
	}

	schemaData, err := s.GithubService.GetRepoFile(ctx, s.PersistentConfig.GitHubToken, repo.Organization, repo.Name, req.SchemaFileSha)
	if err != nil {
		return nil, err
	}

	return &protos.Schema{
		Id:   uuid.NewV4().String(),
		Name: importReq.Name,
		Type: protos.SchemaType_SCHEMA_TYPE_JSONSCHEMA,
		InferredSchema: &encoding.JSONSchemaSettings{
			Schema: schemaData,
		},
		Versions: []*protos.SchemaVersion{
			{
				Version: 1,
				Status:  protos.SchemaStatus_SCHEMA_STATUS_ACCEPTED,
				Files: map[string]string{
					req.SchemaFileName: string(schemaData),
				},
				Settings: &protos.SchemaVersion_JsonSchemaSettings{
					JsonSchemaSettings: &encoding.JSONSchemaSettings{
						Schema: schemaData,
					},
				},
			},
		},
	}, nil
}

// importLocal is used to import a schema from the UI
func (s *Server) importLocal(req *protos.ImportLocalRequest) (*protos.Schema, error) {
	var schema *protos.Schema
	var err error

	switch req.Type {
	case protos.SchemaType_SCHEMA_TYPE_PROTOBUF:
		schema, err = importLocalProtobuf(req)
	case protos.SchemaType_SCHEMA_TYPE_AVRO:
		schema, err = importLocalAvro(req)
	case protos.SchemaType_SCHEMA_TYPE_JSONSCHEMA:
		schema, err = importLocalJSONSchema(req)
	default:
		err = fmt.Errorf("unknown schema type: %d", req.Type)
	}

	if err != nil {
		return nil, err
	}

	return schema, nil
}

// importLocalProtobuf is used to import a protobuf schema from the UI
func importLocalProtobuf(req *protos.ImportLocalRequest) (*protos.Schema, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if len(req.FileContents) == 0 {
		return nil, errors.New("zip archive cannot be empty")
	}

	settings := req.GetProtobufSettings()

	if settings == nil {
		return nil, errors.New("protobuf settings cannot be nil")
	}

	fds, files, err := pb.GetFDFromArchive(req.FileContents, "")
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse protobuf zip")
	}

	mdBlob, err := pb.CreateBlob(fds, settings.ProtobufRootMessage)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create message descriptor blob")
	}

	settings.XMessageDescriptor = mdBlob

	//md, err := pb.GetMDFromDescriptorBlob(mdBlob, req.GetProtobufSettings().ProtobufRootMessage)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to get message descriptor")
	//}

	//inferred, err := jsonschema.InferFromProtobuf("http://plumber.local", "Inferred schema", "", md)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to convert Protobuf schema to JSONSchema")
	//}
	var inferred []byte

	inferredJSON, err := json.Marshal(inferred)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal inferred JSONSchema")
	}

	return &protos.Schema{
		Id:   uuid.NewV4().String(),
		Name: req.Name,
		Type: protos.SchemaType_SCHEMA_TYPE_PROTOBUF,
		InferredSchema: &encoding.JSONSchemaSettings{
			Schema: inferredJSON,
		},
		Versions: []*protos.SchemaVersion{
			{
				Version: 1,
				Status:  protos.SchemaStatus_SCHEMA_STATUS_ACCEPTED,
				Files:   files,
				Settings: &protos.SchemaVersion_ProtobufSettings{
					ProtobufSettings: settings,
				},
			},
		},
	}, nil
}

// importLocalAvro is used to import an avro schema from the UI
func importLocalAvro(req *protos.ImportLocalRequest) (*protos.Schema, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if len(req.FileContents) == 0 {
		return nil, ErrEmptyFile
	}

	//inferred, err := jsonschema.InferFromAvro("http://plumber.local", req.FileContents)
	//if err != nil {
	//	return nil, errors.Wrap(err, "unable to convert Avro to JSONSchema")
	//}
	var inferred []byte

	inferredJSON, err := json.Marshal(inferred)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal inferred JSONSchema")
	}

	return &protos.Schema{
		Id:   uuid.NewV4().String(),
		Name: req.Name,
		Type: protos.SchemaType_SCHEMA_TYPE_AVRO,
		InferredSchema: &encoding.JSONSchemaSettings{
			Schema: inferredJSON,
		},
		Versions: []*protos.SchemaVersion{
			{
				Version: 1,
				Status:  protos.SchemaStatus_SCHEMA_STATUS_ACCEPTED,
				Files: map[string]string{
					req.FileName: string(req.FileContents),
				},
				Settings: &protos.SchemaVersion_AvroSettings{
					AvroSettings: &encoding.AvroSettings{
						Schema: req.FileContents,
					},
				},
			},
		},
	}, nil
}

func importLocalJSONSchema(req *protos.ImportLocalRequest) (*protos.Schema, error) {
	if req == nil {
		return nil, errors.New("request cannot be nil")
	}

	if len(req.FileContents) == 0 {
		return nil, ErrEmptyFile
	}

	return &protos.Schema{
		Id:   uuid.NewV4().String(),
		Name: req.Name,
		Type: protos.SchemaType_SCHEMA_TYPE_JSONSCHEMA,
		InferredSchema: &encoding.JSONSchemaSettings{
			Schema: req.FileContents,
		},
		Versions: []*protos.SchemaVersion{
			{
				Version: 1,
				Status:  protos.SchemaStatus_SCHEMA_STATUS_ACCEPTED,
				Files: map[string]string{
					req.FileName: string(req.FileContents),
				},
				Settings: &protos.SchemaVersion_JsonSchemaSettings{
					JsonSchemaSettings: &encoding.JSONSchemaSettings{
						Schema: req.FileContents,
					},
				},
			},
		},
	}, nil
}
