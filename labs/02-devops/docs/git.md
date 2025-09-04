# Git Command Line Tips

As you begin working with Git and version control, it's helpful to become more familiar with a few key command-line tools. These will make your workflow more efficient and help you understand what Git is doing under the hood.

## Checking repository status
This is one of the most frequently used Git commands. It shows the current state of your working directory and staging area.

Use it to find out:
- Which files have been modified
- Which files are staged for commit
- Which files are untracked

For example:

```bash
git status
```

Might return:

```bash
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)
         modified:   penguins.R
no changes added to commit (use "git add" and/or "git commit -a")
```

This output tells you that you've made changes to `penguins.R`, but haven't yet staged them for commit.

## Adding files
To include changes in your next commit, you need to stage them using git add.

You can specify individual files:

```bash
git add penguins.R other-penguins.R
```

Or stage everything at once with:

```bash
git add .
```

Caution: Be careful when using `git add .`, as it stages all changes, including any you may not have intended to commit.

## Ignoring files
Not all files belong in version control. For example, you may have local settings, credentials, or temporary files that shouldn’t be pushed to GitHub.

To automatically exclude files or patterns, create a `.gitignore` file in your repository directory. List each file or file pattern you want Git to ignore.

Example `.gitignore` file:

```bash
*.log
.DS_Store
credentials.txt
```

Once set, Git will skip tracking those files, even if they exist in your working directory.

## Pulling commits
As you collaborate with others using Git, the remote repository (e.g., on GitHub) may change, especially if your collaborators are actively adding new files or fixing bugs.

To stay up to date with the latest version of the repository, use:

```bash
git pull
```

This command will:
- Contact the remote repository (like GitHub),
- Download any new commits or updates made by others,
- And merge them into your local copy of the project.

You can think of it as asking Git: “Please get all the new stuff my teammates have done and bring it into my local project.”

Under the hood:
git pull is a shortcut for two steps:
1. `git fetch` — grabs the latest changes from the remote but doesn’t apply them yet.
2. `git merge` — integrates those changes into your current branch.

Tip: Always git pull before you start working, to make sure you’re building on the most recent version of the project.

## Cloning repository
Before you can collaborate on a Git project, you need a copy of the repository on your computer. That’s what `git clone` is for.

Use this command to create a local copy of a remote Git repository:

```bash
git clone https://github.com/ucy-coast/cs452-fa25.git
```

This will:
- Download the entire repository, including all files and commit history.
- Set up a link to the remote (usually named origin), so you can pull and push updates.
- Once cloned, you can navigate into the project folder and start working:

```bash
cd cs452-fa25
```

Remember: You only need to git clone once per project. After that, you’ll use git pull to stay in sync.

## Viewing history
Use `git log` to see the history of commits in your repository. Each entry shows:
- The unique commit hash (a long alphanumeric string)
- The author
- The date
- The commit message

Example:

```bash
git log
commit af58f79bfa4301643025dd6c8767e65349cf407a
Author: Name <Email address>
Date:   DD-MM-YYYY
    Add penguin script
```

You can also find this on GitHub, by going to github.com/user-name/repo-name/commits.

You can go back in time to a specific commit, if you know its reference.

## Undoing mistakes
Mistakes happen, even with version control. Fortunately, Git gives you safe ways to undo changes.

If you've committed and pushed something you'd like to reverse, the git revert command is a safe option. It creates a new commit that undoes the effects of a previous one.

Example:

```bash
git revert <hash-of-the-commit-you-want-to-undo>
git push
```

This doesn't delete history. Instead, it records that the change was undone, preserving your project’s full timeline.

Using `git revert` gives you a clear and traceable history. Here’s how it looks in the log:

```bash
git log
commit 6634a076212fb7bac16f9525feae1e83e0f200ca
Author: Name <Email address>
Date:   DD-MM-YYYY
     Revert "Add plain text to code by mistake"
     This reverts commit a8cf7c2592273ef6a28920222a92847794275868.
commit a8cf7c2592273ef6a28920222a92847794275868
Author: Name <Email address>
Date:   DD-MM-YYYY
    Add plain text to code by mistake
```

This shows that the mistaken commit was undone, but both the mistake and the correction are preserved in the project history, a best practice for collaborative and transparent coding.

