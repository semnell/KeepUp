# Contributing to KeepUp

We welcome contributions to KeepUp! By contributing, you can help improve the project and make it even better. Please follow these guidelines to ensure a smooth and effective contribution process.

## Table of Contents

1. [Code of Conduct](#code-of-conduct)
2. [How to Contribute](#how-to-contribute)
3. [Code Style](#code-style)
4. [Testing](#testing)
5. [Pull Request Process](#pull-request-process)

## Code of Conduct

Please read and follow our [Code of Conduct](CODE_OF_CONDUCT.md) to ensure a respectful and inclusive environment for everyone.

## How to Contribute

1. **Fork the Repository:**
   - Click the "Fork" button at the top right of the repository page.

2. **Clone Your Fork:**
   ```sh
   git clone https://github.com/your-username/KeepUp.git
   cd KeepUp
   ```

3. **Create a New Branch:**
   ```sh
   git checkout -b feature/your-feature-name
   ```

4. **Make Your Changes:**
   - Ensure your code follows the project's coding style and conventions.
   - Add tests for your changes if applicable.

5. **Commit Your Changes:**
   ```sh
   git commit -m "Add your commit message here"
   ```

6. **Push Your Changes:**
   ```sh
   git push origin feature/your-feature-name
   ```

7. **Create a Pull Request:**
   - Go to the repository page on GitHub and click the "New Pull Request" button.
   - Provide a clear description of your changes and any related issues.

## Code Style

To ensure consistent code formatting and style across the repository, we use tools like `gofmt` and `revive`. Please run the following commands before committing your changes:

```sh
make fmt
make revive
```

## Testing

Please add tests for your changes and ensure that all tests pass before submitting a pull request. You can run the tests using the following command:

```sh
make test
```

## Pull Request Process

1. Ensure that your changes pass all tests and adhere to the project's coding style.
2. Create a pull request with a clear description of your changes and any related issues.
3. One of the project maintainers will review your pull request and provide feedback if necessary.
4. Once your pull request is approved, it will be merged into the main branch.

Thank you for contributing to KeepUp!
