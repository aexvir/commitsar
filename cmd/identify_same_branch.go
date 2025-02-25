package cmd

import (
	history "github.com/commitsar-app/git/pkg"
)

// IdentifySameBranch breaks up the reference names and tries to identify if the branches are the same
func IdentifySameBranch(branchA, branchB string, gitRepo *history.Git) (bool, error) {
	commitBranchA, err := gitRepo.LatestCommitOnBranch(branchA)

	if err != nil {
		return false, err
	}

	commitBranchB, err := gitRepo.LatestCommitOnBranch(branchB)

	if err != nil {
		return false, err
	}

	return commitBranchA.Hash == commitBranchB.Hash, nil
}
