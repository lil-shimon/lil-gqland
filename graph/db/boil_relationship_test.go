// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import "testing"

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("IssueToUserUsingAuthorUser", testIssueToOneUserUsingAuthorUser)
	t.Run("IssueToRepositoryUsingIssueRepository", testIssueToOneRepositoryUsingIssueRepository)
	t.Run("ProjectcardToPullrequestUsingProjectcardPullrequest", testProjectcardToOnePullrequestUsingProjectcardPullrequest)
	t.Run("ProjectcardToIssueUsingProjectcardIssue", testProjectcardToOneIssueUsingProjectcardIssue)
	t.Run("ProjectcardToProjectUsingProjectcardProject", testProjectcardToOneProjectUsingProjectcardProject)
	t.Run("ProjectToUserUsingOwnerUser", testProjectToOneUserUsingOwnerUser)
	t.Run("PullrequestToRepositoryUsingPullrequestRepository", testPullrequestToOneRepositoryUsingPullrequestRepository)
	t.Run("RepositoryToUserUsingOwnerUser", testRepositoryToOneUserUsingOwnerUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("IssueToProjectcards", testIssueToManyProjectcards)
	t.Run("ProjectToProjectcards", testProjectToManyProjectcards)
	t.Run("PullrequestToProjectcards", testPullrequestToManyProjectcards)
	t.Run("RepositoryToIssues", testRepositoryToManyIssues)
	t.Run("RepositoryToPullrequests", testRepositoryToManyPullrequests)
	t.Run("UserToAuthorIssues", testUserToManyAuthorIssues)
	t.Run("UserToOwnerProjects", testUserToManyOwnerProjects)
	t.Run("UserToOwnerRepositories", testUserToManyOwnerRepositories)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("IssueToUserUsingAuthorIssues", testIssueToOneSetOpUserUsingAuthorUser)
	t.Run("IssueToRepositoryUsingIssues", testIssueToOneSetOpRepositoryUsingIssueRepository)
	t.Run("ProjectcardToPullrequestUsingProjectcards", testProjectcardToOneSetOpPullrequestUsingProjectcardPullrequest)
	t.Run("ProjectcardToIssueUsingProjectcards", testProjectcardToOneSetOpIssueUsingProjectcardIssue)
	t.Run("ProjectcardToProjectUsingProjectcards", testProjectcardToOneSetOpProjectUsingProjectcardProject)
	t.Run("ProjectToUserUsingOwnerProjects", testProjectToOneSetOpUserUsingOwnerUser)
	t.Run("PullrequestToRepositoryUsingPullrequests", testPullrequestToOneSetOpRepositoryUsingPullrequestRepository)
	t.Run("RepositoryToUserUsingOwnerRepositories", testRepositoryToOneSetOpUserUsingOwnerUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {
	t.Run("ProjectcardToPullrequestUsingProjectcards", testProjectcardToOneRemoveOpPullrequestUsingProjectcardPullrequest)
	t.Run("ProjectcardToIssueUsingProjectcards", testProjectcardToOneRemoveOpIssueUsingProjectcardIssue)
}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("IssueToProjectcards", testIssueToManyAddOpProjectcards)
	t.Run("ProjectToProjectcards", testProjectToManyAddOpProjectcards)
	t.Run("PullrequestToProjectcards", testPullrequestToManyAddOpProjectcards)
	t.Run("RepositoryToIssues", testRepositoryToManyAddOpIssues)
	t.Run("RepositoryToPullrequests", testRepositoryToManyAddOpPullrequests)
	t.Run("UserToAuthorIssues", testUserToManyAddOpAuthorIssues)
	t.Run("UserToOwnerProjects", testUserToManyAddOpOwnerProjects)
	t.Run("UserToOwnerRepositories", testUserToManyAddOpOwnerRepositories)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("IssueToProjectcards", testIssueToManySetOpProjectcards)
	t.Run("PullrequestToProjectcards", testPullrequestToManySetOpProjectcards)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("IssueToProjectcards", testIssueToManyRemoveOpProjectcards)
	t.Run("PullrequestToProjectcards", testPullrequestToManyRemoveOpProjectcards)
}
