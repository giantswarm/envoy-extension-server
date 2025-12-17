package extensionserver

import (
	"context"

	pb "github.com/envoyproxy/gateway/proto/extension"
)

func (s *Server) PostTranslateModify(ctx context.Context, req *pb.PostTranslateModifyRequest) (*pb.PostTranslateModifyResponse, error) {
	s.log.Info("PostTranslateModify callback was invoked")

	// Extract CertificatePolicies from the request's extension resources
	policies := s.extractCertificatePolicies(req.PostTranslateContext.GetExtensionResources())

	s.log.Info("fetched CertificatePolicies", "count", len(policies))
	for _, policy := range policies {
		s.log.Info("found CertificatePolicy",
			"name", policy.Name,
			"namespace", policy.Namespace,
			"secretName", policy.Spec.SecretName,
		)
	}

	return &pb.PostTranslateModifyResponse{
		Secrets: req.Secrets,
	}, nil
}
