package readme

import (
	"fmt"

	"github.com/detouri/makemd/internal/markdown"
)

const (
	githubURL     = "https://github.com"
	imgShieldsURL = "https://img.shields.io"
)

func DefaultBadges(cfg ProjectConfig) []string {
	var badges []string
	if cfg.Owner != "" && cfg.Repo != "" {
		switch cfg.CIProvider {
		case "github-actions":
			badges = append(badges,
				markdown.Badge(
					"CI",
					fmt.Sprintf("%s/github/actions/workflow/status/%s/%s/ci.yml?branch=main", imgShieldsURL, cfg.Owner, cfg.Repo),
					fmt.Sprintf("%s/%s/%s/actions/workflows/ci.yml?query=branch%%3Amain", githubURL, cfg.Owner, cfg.Repo),
				),
				markdown.Badge(
					"Issues",
					fmt.Sprintf("%s/github/issues/%s/%s", imgShieldsURL, cfg.Owner, cfg.Repo),
					cfg.IssuesURL,
				),
			)
		default:
			badges = append(badges, markdown.Badge("Build", fmt.Sprintf("%s/badge/build-configure--me-blue", imgShieldsURL)))
		}
	}

	badges = append(badges,
		markdown.Badge("Status", fmt.Sprintf("%s/badge/status-%s-blue", imgShieldsURL, urlEscape(cfg.Status))),
	)

	if !cfg.Private {
		badges = append(badges, markdown.Badge("License", fmt.Sprintf("%s/badge/license-%s-green", imgShieldsURL, urlEscape(cfg.LicenseName))))
	}

	return badges
}
