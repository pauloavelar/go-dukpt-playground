# DUKPT Documentation

This directory contains the GitHub Pages documentation for the DUKPT Playground project.

## Setup

This documentation is designed to be served via GitHub Pages. To enable it:

1. Go to your repository's Settings → Pages
2. Set the source to "Deploy from a branch"
3. Select `main` branch and `/docs` folder
4. The documentation will be available at `https://[username].github.io/[repository-name]/`

## Structure

- `index.html` - Main documentation page with complete DUKPT workflow
- `assets/css/style.css` - Clean, responsive styling
- Text-based ASCII diagrams for compatibility and clarity

## Features

- **Complete DUKPT Documentation**: Covers the entire workflow from BDK to encrypted data
- **Interactive Navigation**: Smooth scrolling between sections
- **Text-based Diagrams**: ASCII art diagrams that work without external dependencies
- **Responsive Design**: Works on desktop and mobile devices
- **Clean Styling**: Professional appearance following community standards
- **Real Examples**: Includes actual examples from the Go implementation

## Diagram Options

The documentation uses ASCII art diagrams for maximum compatibility and to meet the requirement for text-based diagrams. This approach:

- Works without external CDN dependencies
- Loads instantly on GitHub Pages
- Is accessible and printer-friendly
- Follows the "text-based as much as possible" requirement

Alternative diagram options that could be integrated in the future:
- **Mermaid.js**: For more complex sequence diagrams (requires CDN)
- **PlantUML**: Would need server-side rendering or JavaScript library
- **SVG**: For custom graphics (would need to be created manually)

## Content Overview

1. **DUKPT Overview**: Introduction and key concepts
2. **Complete Workflow**: Full transaction flow diagram
3. **IPEK Derivation**: Detailed process with step-by-step explanation
4. **Session Key Derivation**: NRKGP algorithm and implementation
5. **Data Encryption**: Track2 formatting and encryption process
6. **Practical Examples**: Real code examples and test data

The documentation is designed to be educational and comprehensive while maintaining clean, professional styling.