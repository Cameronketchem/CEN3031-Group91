# Contributing
Workflow & guide for contributing to the CEN3031-Group91 project.

## Contribution Workflow
1. (Optional) fork the project repository.
2. Checkout a new branch from `main`.
   - If on a fork, sync your fork to the latest upstream version before branching.
   - ALWAYS branch from the latest version of `main`.
   - Each PR/feature goes on its own branch (again, branched from `main`). This ensures that new features are not dependent on uncommitted code.
3. Commit changes
   - Refer to [Commit Messages](#commit-messages) for writing commit messages.
   - Write a separate commit message for every file that you've modified, unless the same message applies to multiple files (in which case they can be packed into a single commit).
4. Write unit tests (if applicable).
5. Push to origin with `git push origin [your branch name]` on your terminal.
6. Create a pull request.
   - If you've made changes to a fork, make sure you're opening a pull request to the original remote repository.
   - Always open PR's on the `main` branch unless otherwise instructed
   - Add a descriptive title and describe your changes.
   - Refer to [Pull Requests](#pull-requests)

## Commit Messages
Commit messages should be written for each file that's been modified. Avoid generic messages like "made changes." 
Commit messages are preceded by a prefix describing the category of the commit. See [Message Prefixes](#message-prefixes).

### Writing messages
That said, avoid commiting large changes all at once: messages like `updated app` or `fixed problem` should be broken down into their specific constituents. For example:
1. `fixed problem` should ideally be broken up into multiple commits like `fix: added missing form validation` or `fix: typo`.
2. `updated app`: What exactly did you update and how did you do it? Break up each change into a commit like `feat: added form element for xyz`, or `feat: added password field`.

### Message Prefixes
- `feat: ...`: For new features.
- `docs: ...`: For adding or changing documentation
- `ci: ...`: For modifying CI, jobs or workflows
- `fix: ...`: For when you've fixed something.
- `lint: ...`: For when you've fixed/changed code syntax.
- `refactor: ...`: For when you've rewritten or reorganized something. Use this if you've moved files or created folders. This is different from `lint` because it doesn't necessarily describe syntactical changes, it describes organizational changes.
- `deps: ...`: When installing or updating dependencies. Package-lock.json changes will typically use this.

