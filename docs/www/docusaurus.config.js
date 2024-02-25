import { themes as prismThemes } from 'prism-react-renderer';

const config = {
  title: 'Kurama',
  tagline: 'A simple, blazingly fast, cloud-native job scheduler and orchestrator written in Go.',
  favicon: 'img/favicon.png',
  url: 'https://iwas-coder.github.io',
  baseUrl: '/kurama',
  organizationName: 'iWas-Coder',
  projectName: 'kurama',
  trailingSlash: false,
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',

  i18n: {
    defaultLocale: 'en',
    locales: ['en']
  },

  presets: [
    [
      'classic',
      {
        docs: {
          sidebarPath: './sidebars.js',
          editUrl: 'https://github.com/iWas-Coder/kurama/docs/www/docs/',
          showLastUpdateTime: true
        },
        blog: {
          showReadingTime: true,
          editUrl: 'https://github.com/iWas-Coder/kurama/docs/www/blog/'
        },
        theme: {
          customCss: './src/css/custom.css',
        }
      }
    ]
  ],

  themeConfig: ({
    image: 'img/docusaurus-social-card.jpg',
    navbar: {
      logo: {
        alt: 'Kurama Logo',
        src: 'img/logo.png'
      },
      items: [
        {
          href: 'https://github.com/iWas-Coder/kurama',
          label: 'GitHub',
          position: 'left',
        },
        {
          type: 'docSidebar',
          sidebarId: 'docs',
          position: 'left',
          label: 'Docs',
        },
        {
          to: 'blog',
          position: 'left',
          label: 'Blog'
        }
      ],
    },
    agolia: {
      appId: '',
      apiKey: '',
      indexName: '',
      contextualSearch: true
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'What is Kurama?',
              to: 'docs/overview/1.1-what-is-kurama'
            },
            {
              label: 'Getting started',
              to: 'docs/getting-started/2.1-getting-started'
            }
          ]
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Contributing',
              href: 'https://github.com/iWas-Coder/kurama/blob/master/CONTRIBUTING.org'
            }
          ]
        },
        {
          title: 'More',
          items: [
            {
              label: 'GitHub',
              href: 'https://github.com/iWas-Coder/kurama',
            }
          ]
        }
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Wasym Atieh Alonso. All rights reserved.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
    }
  })
};

export default config;
