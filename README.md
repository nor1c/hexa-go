### Hexagonal Architecture

## Directory Structure
- `pkg/http` - contains handler for modules, think it like a controller in MVC
- `pkg/domain` - contains models for the modules
- `pkg/ports` - contains interfaces for `repositories` and `usecases`