# Solutions Engineer Work Sample Test

This is the work sample test for our solutions engineer role. We ask you to
complete this project so we can assess you for the work we expect you to be
doing at RainforestQA.

We believe QA should be done as often and as soon as possible, which means
ideally Rainforest is configured as part of customer's CI/CD pipeline. We want
you to demonstrate your ability to configure a CI/CD pipeline by setting up a
CI/CD pipeline for this repo.

The pipeline will:
* Build the code in this repo
* Run the unit tests for the code
* Deploy the code to a "QA" environment
* Run a Rainforest test against the QA environment
* Deploy the code to a "production" environment

The pipeline should run whenever changes are merged into the `main` branch of
the repo.

You may choose whichever CI/CD and hosting provider you are most confortable
with. We recommend Github actions and fly.io, as they both have free tiers
which are more than adaquate for completing this task.

You can signup for a Rainforest account at https://app.rainforestqa.com/auth/signup.
The Rainforest test only has to check that the site is up and the logo and
banner are displayed.

You should take no more than 3 hours to complete this task.

To complete this task and submit your work:
* Create a **private** repository from this repo [using the template mechanism](https://docs.github.com/en/repositories/creating-and-managing-repositories/creating-a-repository-from-a-template)
* Make the changes required to implement the CI/CD pipeline
* Let us know that you've finished the work
* We will ask you to [add the interviewer as a collaborator](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-access-to-your-personal-repositories/inviting-collaborators-to-a-personal-repository) on the repo so they can review the work
