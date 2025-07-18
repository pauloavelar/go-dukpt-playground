# DUKPT Workflow Documentation

This directory contains the GitHub Pages documentation for the DUKPT (Derived Unique Key Per Transaction) workflow implementation.

## Overview

The documentation provides a comprehensive guide to understanding DUKPT key management in payment processing systems, featuring:

- Complete workflow explanations
- Interactive Mermaid.js diagrams
- Real code examples from the Go implementation
- Professional, responsive design

## Features

### Modern Diagrams
The documentation uses **Mermaid.js diagrams** loaded from an external CDN for professional visualization:
- Sequence diagrams for transaction flows
- Flowcharts for key derivation processes
- Clean, modern styling with neutral theme

### Comprehensive Coverage
- DUKPT Overview and Security Benefits
- Complete Transaction Flow
- IPEK Derivation Process  
- Session Key Derivation (NRKGP)
- Data Encryption Process
- Practical Examples

## Setup for GitHub Pages

### Automatic Deployment

GitHub Pages is automatically deployed via GitHub Actions when changes are pushed to the `main` branch. The documentation is available at `https://pauloavelar.github.io/go-dukpt-playground/`.

### Preview Deployments

Pull requests with changes to the documentation automatically generate preview deployments with preview URLs posted as comments on pull requests.

### Manual Setup (if needed)

To manually enable GitHub Pages:

1. Go to repository **Settings → Pages**  
2. Set source to **"GitHub Actions"**
3. The deployment workflow will handle the rest

### Local Development

For local development and testing:

```bash
# From repository root
./preview-docs.sh
```

This script starts a local HTTP server at `http://localhost:8080` for previewing documentation changes before committing.

## Technical Details

### Dependencies
- **Mermaid.js v10.6.1**: External CDN for diagram rendering
- **No build process required**: Self-contained HTML/CSS
- **Responsive design**: Works on desktop and mobile

### File Structure
```
docs/
├── index.html          # Main documentation
├── assets/
│   └── css/
│       └── style.css   # Custom styling
└── README.md          # This file
```

### Browser Compatibility
The documentation works in all modern browsers that support:
- ES6 JavaScript features
- CSS Grid and Flexbox
- External script loading

### CDN Usage
As requested, the documentation uses Mermaid.js from an external CDN:
```html
<script src="https://cdn.jsdelivr.net/npm/mermaid@10.6.1/dist/mermaid.min.js"></script>
```

If the CDN is unavailable, diagrams will display as readable text with a notice to users.

## Content

The documentation covers the complete DUKPT workflow with detailed explanations of:

1. **Key Components**: BDK, KSN, IPEK, Session Keys
2. **Security Benefits**: Unique keys per transaction, non-reversible process
3. **Implementation Details**: Step-by-step algorithm explanations
4. **Real Examples**: Actual test data and results

All technical information is backed by the actual Go implementation in this repository.