# Docker Environment Configuration Guide

## Environment Structure

The Docker configuration is separated into three distinct environments:

```
docker/
├── local/         # For local development with hot-reloading
├── ci/tests/      # For running automated tests in CI
└── cloud/         # For deployments
```

## Using the Environments

### Local Development

For local development with hot-reloading:

```bash
make docker-local
```

This will:
- Mount your source code as a volume for live changes
- Set up hot-reloading with reflex
- Expose ports 8080 (API) and 8081 (Smocker UI)

To run in detached mode:

```bash
make docker-local-detached
```

### CI/Testing

For running tests in CI:

```bash
# Run all tests (unit + integration)
make docker-test

# Run only unit tests (faster)
make docker-test-unit

# Run only linting
make docker-test-lint

# Run the full CI pipeline
make docker-ci
```

The CI configuration:
- Skips the frontend build completely
- Focuses only on backend testing
- Has configurable test types via environment variables

### Cloud

For building a usage-ready image:

```bash
# Build with version info
make docker-cloud VERSION=1.2.3 COMMIT=$(git rev-parse HEAD)

# Run the container
make docker-cloud-run
```

This creates an optimized image for deployment with:
- Multi-stage build for minimal image size
- Properly tagged with version information

## .dockerignore Files

Each environment has its own `.dockerignore` file that extends the base configuration:

```
.cache/
.git/
.vscode/
node_modules/
tests/
```

- **Local**: Excludes build artifacts and CI test files
- **CI/Tests**: Excludes frontend code and dependencies
- **Cloud**: Comprehensive exclusions for a minimal image

## Additional Information

- All Docker configurations use the same base code
- Environment-specific optimizations are applied to each Dockerfile
- The base `.dockerignore` is preserved and extended for each environment
