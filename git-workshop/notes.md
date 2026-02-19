# GitHub CLI

To install CLI tool:
`brew install gh`

*login with GitHub credentials*

`gh repo create`

# .gitignore
- Don't share your secrets.

# git blame from GitHub
- See who wrote what code in your IDE.
  - Use Gitlens VSCode extension.

# Cloning from remote
- Don't push large binaries!
  - Git commit object model:
    - Commits -> trees -> blobs.
      - Each new commit is a snapshot of the repository.
      - Essentially 'diffs' on top of each other.
      - Each tree is a list of all the files in the repo.
        - Trees point to BLOBs which are the data files.
        - BLOBs only are duplicated if the file changed.
          - Text files compress really well in diffs, but binaries change a lot.
    - if a 500MB binary is modified 20 times, it could be 10GB.
    - Even old commits.
  - See my failure.

# Pull Requests
- Create pull requests for others to review your code.

- Create *draft* pull requests to put up your code for others to see, but still WIP.

# Reviewing in GitHub UI

# Issues
- Make tickets and assign them.
- Move on ticket board.

# Protecting main
- No force pushes to main.
- Require a pull request to merge into main.
  - Can bypass this for admins.

# GitHub Actions
- YML files in .github/workflows
  - Configurable automated process.
    - Can define workflows to check for formatting on code.
    - Can define a job to deploy your code to a server when it is pushed.
    - Can run tests on push.
      https://github.com/GenerateNU/prisere/commits/main/

=====
name: Hello Demo

on:
  push:
    branches: \[main]

jobs:
  greet:
    runs-on: ubuntu-latest
    steps:
      - name: Say hello
        run: echo "Hello from GitHub Actions!"
=====
