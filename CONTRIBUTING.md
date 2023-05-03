# **Contribution Guide for OP Vault Project**

This contribution guide will help you understand how to contribute to the OP Vault project. It will walk you through the process of setting up the development environment, creating new features or fixing bugs, and submitting a pull request.

## **Table of Contents**

- [Prerequisites](#prerequisites)
- [Fork and Clone the Repository](#fork-and-clone-the-repository)
- [Set Up the Development Environment](#set-up-the-development-environment)
- [Create a New Branch](#create-a-new-branch)
- [Work on New Features or Bug Fixes](#work-on-new-features-or-bug-fixes)
- [Test Your Changes](#test-your-changes)
- [Commit and Push Your Changes](#commit-and-push-your-changes)
- [Create a Pull Request](#create-a-new-branch)
- [Monitor Your Pull Request](#monitor-your-pull-request)

## **Prerequisites**

Before you start contributing to the OP Vault project, ensure you have the following installed on your system:

- [Git](https://git-scm.com/downloads)
- [Go](https://go.dev/dl/)
- [Node.js](https://nodejs.org/en/download)
- [Poppler](https://poppler.freedesktop.org/)

Also, make sure you have an account on GitHub and OpenAI.

## **Fork and Clone the Repository**

- Go to the OP Vault repository on GitHub.
- Click the "Fork" button in the top-right corner to create a copy of the repository under your GitHub account.
- Open a terminal or command prompt on your computer.
- Clone the forked repository to your local machine by running the following command (replace your_username with your GitHub username):

      git clone https://github.com/your_username/vault-ai.git

## **Set Up the Development Environment**

- Change the directory to the cloned repository:

      cd vault-ai

- Create the necessary files for storing API keys and endpoints in the secret folder:

      echo "your_openai_api_key_here" > secret/openai_api_key
      echo "your_pinecone_api_key_here" > secret/pinecone_api_key
      echo "https://example-50709b5.svc.asia-southeast1-gcp.pinecone.io" > secret/pinecone_api_endpoint

- Install the required JavaScript dependencies:

      npm install

## **Create a New Branch**

Before you start working on a new feature or bug fix, create a new branch to keep your changes separate from the main branch. Name the branch according to the feature or bug fix you are working on.

    git checkout -b your-feature-branch

## **Work on New Features or Bug Fixes**

- Open your favorite code editor and start working on the new feature or bug fix.
- Make sure to follow the coding style and conventions of the project.
- Write clean, efficient, and well-documented code.

## **Test Your Changes**

- Start the Golang web server:

      npm start

- In another terminal window, run Webpack to compile the JavaScript code and create a bundle.js file:

      npm run dev

- Visit the local version of the site at http://localhost:8100 and test your changes.

## **Commit and Push Your Changes**

- Stage your changes:

      git add .

- Commit your changes with a meaningful commit message:

      git commit -m "Your commit message here"

- Push your changes to your forked repository on GitHub:

      git push origin your-feature-branch

## **Create a Pull Request**

- Go to your forked repository on GitHub.
- Click on the "Pull requests" tab and click the "New pull request" button.
- Select the pashpashpash/vault-ai repository as the base repository and your forked repository as the head repository.
- Choose the branch you created earlier (your-feature-branch) as the head branch.
- Write a clear and concise title and description for your pull request.
- Click the "Create pull request" button.

## **Monitor Your Pull Request**

- Keep an eye on your pull request for any feedback, comments, or requests for changes from the project maintainers.
- If changes are requested, make the necessary changes in your local branch, commit them, and push the updates to your forked repository.
- The pull request will be updated automatically, and the project maintainers will review your changes again.

Congratulations! You have successfully contributed to the OP Vault project. Keep contributing and improving the project for the benefit of the community.