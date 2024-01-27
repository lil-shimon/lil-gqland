// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testProjectcards(t *testing.T) {
	t.Parallel()

	query := Projectcards()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProjectcardsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectcardsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Projectcards().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectcardsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProjectcardSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProjectcardsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProjectcardExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Projectcard exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProjectcardExists to return true, but got false.")
	}
}

func testProjectcardsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	projectcardFound, err := FindProjectcard(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if projectcardFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProjectcardsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Projectcards().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProjectcardsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Projectcards().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProjectcardsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	projectcardOne := &Projectcard{}
	projectcardTwo := &Projectcard{}
	if err = randomize.Struct(seed, projectcardOne, projectcardDBTypes, false, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}
	if err = randomize.Struct(seed, projectcardTwo, projectcardDBTypes, false, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = projectcardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = projectcardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Projectcards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProjectcardsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	projectcardOne := &Projectcard{}
	projectcardTwo := &Projectcard{}
	if err = randomize.Struct(seed, projectcardOne, projectcardDBTypes, false, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}
	if err = randomize.Struct(seed, projectcardTwo, projectcardDBTypes, false, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = projectcardOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = projectcardTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func projectcardBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func projectcardAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Projectcard) error {
	*o = Projectcard{}
	return nil
}

func testProjectcardsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Projectcard{}
	o := &Projectcard{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, projectcardDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Projectcard object: %s", err)
	}

	AddProjectcardHook(boil.BeforeInsertHook, projectcardBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	projectcardBeforeInsertHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.AfterInsertHook, projectcardAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	projectcardAfterInsertHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.AfterSelectHook, projectcardAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	projectcardAfterSelectHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.BeforeUpdateHook, projectcardBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	projectcardBeforeUpdateHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.AfterUpdateHook, projectcardAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	projectcardAfterUpdateHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.BeforeDeleteHook, projectcardBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	projectcardBeforeDeleteHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.AfterDeleteHook, projectcardAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	projectcardAfterDeleteHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.BeforeUpsertHook, projectcardBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	projectcardBeforeUpsertHooks = []ProjectcardHook{}

	AddProjectcardHook(boil.AfterUpsertHook, projectcardAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	projectcardAfterUpsertHooks = []ProjectcardHook{}
}

func testProjectcardsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProjectcardsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(projectcardColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProjectcardToOnePullrequestUsingProjectcardPullrequest(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Projectcard
	var foreign Pullrequest

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, pullrequestDBTypes, false, pullrequestColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pullrequest struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.Pullrequest, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ProjectcardPullrequest().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddPullrequestHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Pullrequest) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ProjectcardSlice{&local}
	if err = local.L.LoadProjectcardPullrequest(ctx, tx, false, (*[]*Projectcard)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ProjectcardPullrequest == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ProjectcardPullrequest = nil
	if err = local.L.LoadProjectcardPullrequest(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ProjectcardPullrequest == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testProjectcardToOneIssueUsingProjectcardIssue(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Projectcard
	var foreign Issue

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, issueDBTypes, false, issueColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Issue struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.Issue, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ProjectcardIssue().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddIssueHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Issue) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ProjectcardSlice{&local}
	if err = local.L.LoadProjectcardIssue(ctx, tx, false, (*[]*Projectcard)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ProjectcardIssue == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ProjectcardIssue = nil
	if err = local.L.LoadProjectcardIssue(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ProjectcardIssue == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testProjectcardToOneProjectUsingProjectcardProject(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Projectcard
	var foreign Project

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, projectcardDBTypes, false, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, projectDBTypes, false, projectColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Project struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.Project = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ProjectcardProject().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddProjectHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Project) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := ProjectcardSlice{&local}
	if err = local.L.LoadProjectcardProject(ctx, tx, false, (*[]*Projectcard)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ProjectcardProject == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ProjectcardProject = nil
	if err = local.L.LoadProjectcardProject(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ProjectcardProject == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testProjectcardToOneSetOpPullrequestUsingProjectcardPullrequest(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Projectcard
	var b, c Pullrequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectcardDBTypes, false, strmangle.SetComplement(projectcardPrimaryKeyColumns, projectcardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pullrequestDBTypes, false, strmangle.SetComplement(pullrequestPrimaryKeyColumns, pullrequestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, pullrequestDBTypes, false, strmangle.SetComplement(pullrequestPrimaryKeyColumns, pullrequestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Pullrequest{&b, &c} {
		err = a.SetProjectcardPullrequest(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ProjectcardPullrequest != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Projectcards[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.Pullrequest, x.ID) {
			t.Error("foreign key was wrong value", a.Pullrequest)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Pullrequest))
		reflect.Indirect(reflect.ValueOf(&a.Pullrequest)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.Pullrequest, x.ID) {
			t.Error("foreign key was wrong value", a.Pullrequest, x.ID)
		}
	}
}

func testProjectcardToOneRemoveOpPullrequestUsingProjectcardPullrequest(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Projectcard
	var b Pullrequest

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectcardDBTypes, false, strmangle.SetComplement(projectcardPrimaryKeyColumns, projectcardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, pullrequestDBTypes, false, strmangle.SetComplement(pullrequestPrimaryKeyColumns, pullrequestColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetProjectcardPullrequest(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveProjectcardPullrequest(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ProjectcardPullrequest().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ProjectcardPullrequest != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.Pullrequest) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Projectcards) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testProjectcardToOneSetOpIssueUsingProjectcardIssue(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Projectcard
	var b, c Issue

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectcardDBTypes, false, strmangle.SetComplement(projectcardPrimaryKeyColumns, projectcardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Issue{&b, &c} {
		err = a.SetProjectcardIssue(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ProjectcardIssue != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Projectcards[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.Issue, x.ID) {
			t.Error("foreign key was wrong value", a.Issue)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Issue))
		reflect.Indirect(reflect.ValueOf(&a.Issue)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.Issue, x.ID) {
			t.Error("foreign key was wrong value", a.Issue, x.ID)
		}
	}
}

func testProjectcardToOneRemoveOpIssueUsingProjectcardIssue(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Projectcard
	var b Issue

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectcardDBTypes, false, strmangle.SetComplement(projectcardPrimaryKeyColumns, projectcardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, issueDBTypes, false, strmangle.SetComplement(issuePrimaryKeyColumns, issueColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetProjectcardIssue(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveProjectcardIssue(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.ProjectcardIssue().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.ProjectcardIssue != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.Issue) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.Projectcards) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testProjectcardToOneSetOpProjectUsingProjectcardProject(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Projectcard
	var b, c Project

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, projectcardDBTypes, false, strmangle.SetComplement(projectcardPrimaryKeyColumns, projectcardColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, projectDBTypes, false, strmangle.SetComplement(projectPrimaryKeyColumns, projectColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Project{&b, &c} {
		err = a.SetProjectcardProject(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ProjectcardProject != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Projectcards[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Project != x.ID {
			t.Error("foreign key was wrong value", a.Project)
		}

		zero := reflect.Zero(reflect.TypeOf(a.Project))
		reflect.Indirect(reflect.ValueOf(&a.Project)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.Project != x.ID {
			t.Error("foreign key was wrong value", a.Project, x.ID)
		}
	}
}

func testProjectcardsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProjectcardsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProjectcardSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProjectcardsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Projectcards().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	projectcardDBTypes = map[string]string{`ID`: `TEXT`, `Project`: `TEXT`, `Issue`: `TEXT`, `Pullrequest`: `TEXT`}
	_                  = bytes.MinRead
)

func testProjectcardsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(projectcardPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(projectcardAllColumns) == len(projectcardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProjectcardsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(projectcardAllColumns) == len(projectcardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Projectcard{}
	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, projectcardDBTypes, true, projectcardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(projectcardAllColumns, projectcardPrimaryKeyColumns) {
		fields = projectcardAllColumns
	} else {
		fields = strmangle.SetComplement(
			projectcardAllColumns,
			projectcardPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := ProjectcardSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProjectcardsUpsert(t *testing.T) {
	t.Parallel()
	if len(projectcardAllColumns) == len(projectcardPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Projectcard{}
	if err = randomize.Struct(seed, &o, projectcardDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Projectcard: %s", err)
	}

	count, err := Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, projectcardDBTypes, false, projectcardPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Projectcard struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Projectcard: %s", err)
	}

	count, err = Projectcards().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
