# validate-a-changelog

Parse and validate Changelog following [keep a changelog](https://keepachangelog.com/en/1.1.0/) convention.

## cmd/parse-changelog

```
Usage: parse-changelog <file> [version]
```

Sample output for the test changelog in keep a changelog website:

```json
{
  "title": "Changelog",
  "versions": [
    {
      "version": "Unreleased",
      "release_date": null,
      "entries": {
        "Added": [
          {
            "description": "v1.1 Brazilian Portuguese translation."
          },
          {
            "description": "v1.1 German Translation"
          },
          {
            "description": "v1.1 Spanish translation."
          },
          {
            "description": "v1.1 Italian translation."
          },
          {
            "description": "v1.1 Polish translation."
          },
          {
            "description": "v1.1 Ukrainian translation."
          }
        ],
        "Change": [
          {
            "description": "Use frontmatter title & description in each language version template"
          },
          {
            "description": "Replace broken OpenGraph image with an appropriately-sized Keep a Changelog "
          },
          {
            "description": "Fix OpenGraph title & description for all languages so the title and "
          }
        ],
        "Removed": [
          {
            "description": "Trademark sign previously shown after the project description in version "
          }
        ]
      }
    },
    {
      "version": "1.1.1",
      "release_date": "2023-03-05T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "Arabic translation (#444)."
          },
          {
            "description": "v1.1 French translation."
          },
          {
            "description": "v1.1 Dutch translation (#371)."
          },
          {
            "description": "v1.1 Russian translation (#410)."
          },
          {
            "description": "v1.1 Japanese translation (#363)."
          },
          {
            "description": "v1.1 Norwegian Bokmål translation (#383)."
          },
          {
            "description": "v1.1 \"Inconsistent Changes\" Turkish translation (#347)."
          },
          {
            "description": "Default to most recent versions available for each languages."
          },
          {
            "description": "Display count of available translations (26 to date!)."
          },
          {
            "description": "Centralize all links into `/data/links.json` so they can be updated easily."
          }
        ],
        "Changed": [
          {
            "description": "Upgrade dependencies: Ruby 3.2.1, Middleman, etc."
          }
        ],
        "Fixed": [
          {
            "description": "Improve French translation (#377)."
          },
          {
            "description": "Improve id-ID translation (#416)."
          },
          {
            "description": "Improve Persian translation (#457)."
          },
          {
            "description": "Improve Russian translation (#408)."
          },
          {
            "description": "Improve Swedish title (#419)."
          },
          {
            "description": "Improve zh-CN translation (#359)."
          },
          {
            "description": "Improve French translation (#357)."
          },
          {
            "description": "Improve zh-TW translation (#360, #355)."
          },
          {
            "description": "Improve Spanish (es-ES) transltion (#362)."
          },
          {
            "description": "Foldout menu in Dutch translation (#371)."
          },
          {
            "description": "Missing periods at the end of each change (#451)."
          },
          {
            "description": "Fix missing logo in 1.1 pages."
          },
          {
            "description": "Display notice when translation isn't for most recent version."
          },
          {
            "description": "Various broken links, page versions, and indentations."
          }
        ],
        "Removed": [
          {
            "description": "Unused normalize.css file."
          },
          {
            "description": "Identical links assigned in each translation file."
          },
          {
            "description": "Duplicate index file for the english version."
          }
        ]
      }
    },
    {
      "version": "1.1.0",
      "release_date": "2019-02-15T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "Danish translation (#297)."
          },
          {
            "description": "Georgian translation from (#337)."
          },
          {
            "description": "Changelog inconsistency section in Bad Practices."
          }
        ],
        "Fixed": [
          {
            "description": "Italian translation (#332)."
          },
          {
            "description": "Indonesian translation (#336)."
          }
        ]
      }
    },
    {
      "version": "1.0.0",
      "release_date": "2017-06-20T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "New visual identity by [@tylerfortune8](https://github.com/tylerfortune8)."
          },
          {
            "description": "Version navigation."
          },
          {
            "description": "Links to latest released version in previous versions."
          },
          {
            "description": "\"Why keep a changelog?\" section."
          },
          {
            "description": "\"Who needs a changelog?\" section."
          },
          {
            "description": "\"How do I make a changelog?\" section."
          },
          {
            "description": "\"Frequently Asked Questions\" section."
          },
          {
            "description": "New \"Guiding Principles\" sub-section to \"How do I make a changelog?\"."
          },
          {
            "description": "Simplified and Traditional Chinese translations from [@tianshuo](https://github.com/tianshuo)."
          },
          {
            "description": "German translation from [@mpbzh](https://github.com/mpbzh) & [@Art4](https://github.com/Art4)."
          },
          {
            "description": "Italian translation from [@azkidenz](https://github.com/azkidenz)."
          },
          {
            "description": "Swedish translation from [@magol](https://github.com/magol)."
          },
          {
            "description": "Turkish translation from [@emreerkan](https://github.com/emreerkan)."
          },
          {
            "description": "French translation from [@zapashcanon](https://github.com/zapashcanon)."
          },
          {
            "description": "Brazilian Portuguese translation from [@Webysther](https://github.com/Webysther)."
          },
          {
            "description": "Polish translation from [@amielucha](https://github.com/amielucha) & [@m-aciek](https://github.com/m-aciek)."
          },
          {
            "description": "Russian translation from [@aishek](https://github.com/aishek)."
          },
          {
            "description": "Czech translation from [@h4vry](https://github.com/h4vry)."
          },
          {
            "description": "Slovak translation from [@jkostolansky](https://github.com/jkostolansky)."
          },
          {
            "description": "Korean translation from [@pierceh89](https://github.com/pierceh89)."
          },
          {
            "description": "Croatian translation from [@porx](https://github.com/porx)."
          },
          {
            "description": "Persian translation from [@Hameds](https://github.com/Hameds)."
          },
          {
            "description": "Ukrainian translation from [@osadchyi-s](https://github.com/osadchyi-s)."
          }
        ],
        "Changed": [
          {
            "description": "Start using \"changelog\" over \"change log\" since it's the common usage."
          },
          {
            "description": "Start versioning based on the current English version at 0.3.0 to help"
          },
          {
            "description": "Rewrite \"What makes unicorns cry?\" section."
          },
          {
            "description": "Rewrite \"Ignoring Deprecations\" sub-section to clarify the ideal"
          },
          {
            "description": "Improve \"Commit log diffs\" sub-section to further argument against"
          },
          {
            "description": "Merge \"Why can’t people just use a git log diff?\" with \"Commit log"
          },
          {
            "description": "Fix typos in Simplified Chinese and Traditional Chinese translations."
          },
          {
            "description": "Fix typos in Brazilian Portuguese translation."
          },
          {
            "description": "Fix typos in Turkish translation."
          },
          {
            "description": "Fix typos in Czech translation."
          },
          {
            "description": "Fix typos in Swedish translation."
          },
          {
            "description": "Improve phrasing in French translation."
          },
          {
            "description": "Fix phrasing and spelling in German translation."
          }
        ],
        "Removed": [
          {
            "description": "Section about \"changelog\" vs \"CHANGELOG\"."
          }
        ]
      }
    },
    {
      "version": "0.3.0",
      "release_date": "2015-12-03T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "RU translation from [@aishek](https://github.com/aishek)."
          },
          {
            "description": "pt-BR translation from [@tallesl](https://github.com/tallesl)."
          },
          {
            "description": "es-ES translation from [@ZeliosAriex](https://github.com/ZeliosAriex)."
          }
        ]
      }
    },
    {
      "version": "0.2.0",
      "release_date": "2015-10-06T00:00:00Z",
      "entries": {
        "Changed": [
          {
            "description": "Remove exclusionary mentions of \"open source\" since this project can"
          }
        ]
      }
    },
    {
      "version": "0.1.0",
      "release_date": "2015-10-06T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "Answer \"Should you ever rewrite a change log?\"."
          }
        ],
        "Changed": [
          {
            "description": "Improve argument against commit logs."
          },
          {
            "description": "Start following [SemVer](https://semver.org) properly."
          }
        ]
      }
    },
    {
      "version": "0.0.8",
      "release_date": "2015-02-17T00:00:00Z",
      "entries": {
        "Changed": [
          {
            "description": "Update year to match in every README example."
          },
          {
            "description": "Reluctantly stop making fun of Brits only, since most of the world"
          }
        ],
        "Fixed": [
          {
            "description": "Fix typos in recent README changes."
          },
          {
            "description": "Update outdated unreleased diff link."
          }
        ]
      }
    },
    {
      "version": "0.0.7",
      "release_date": "2015-02-16T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "Link, and make it obvious that date format is ISO 8601."
          }
        ],
        "Changed": [
          {
            "description": "Clarified the section on \"Is there a standard change log format?\"."
          }
        ],
        "Fixed": [
          {
            "description": "Fix Markdown links to tag comparison URL with footnote-style links."
          }
        ]
      }
    },
    {
      "version": "0.0.6",
      "release_date": "2014-12-12T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "README section on \"yanked\" releases."
          }
        ]
      }
    },
    {
      "version": "0.0.5",
      "release_date": "2014-08-09T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "Markdown links to version tags on release headings."
          },
          {
            "description": "Unreleased section to gather unreleased changes and encourage note"
          }
        ]
      }
    },
    {
      "version": "0.0.4",
      "release_date": "2014-08-09T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "Better explanation of the difference between the file (\"CHANGELOG\")"
          }
        ],
        "Changed": [
          {
            "description": "Refer to a \"change log\" instead of a \"CHANGELOG\" throughout the site"
          }
        ],
        "Removed": [
          {
            "description": "Remove empty sections from CHANGELOG, they occupy too much space and"
          }
        ]
      }
    },
    {
      "version": "0.0.3",
      "release_date": "2014-08-09T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "\"Why should I care?\" section mentioning The Changelog podcast."
          }
        ]
      }
    },
    {
      "version": "0.0.2",
      "release_date": "2014-07-10T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "Explanation of the recommended reverse chronological release ordering."
          }
        ]
      }
    },
    {
      "version": "0.0.1",
      "release_date": "2014-05-31T00:00:00Z",
      "entries": {
        "Added": [
          {
            "description": "This CHANGELOG file to hopefully serve as an evolving example of a"
          },
          {
            "description": "CNAME file to enable GitHub Pages custom domain."
          },
          {
            "description": "README now contains answers to common questions about CHANGELOGs."
          },
          {
            "description": "Good examples and basic guidelines, including proper date formatting."
          },
          {
            "description": "Counter-examples: \"What makes unicorns cry?\"."
          }
        ]
      }
    }
  ]
}
```

## cmd/validate-changelog

```
Usage: validate-changelog [-allow-empty-version] [-allow-invalid-change-type] [-allow-missing-release-date] [-json] <file>
```

## cmd/lint-changelog

```
Usage: lint-changelog [-json] <file>
```