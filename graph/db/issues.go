// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Issue is an object representing the database table.
type Issue struct {
	ID         string `boil:"id" json:"id" toml:"id" yaml:"id"`
	URL        string `boil:"url" json:"url" toml:"url" yaml:"url"`
	Title      string `boil:"title" json:"title" toml:"title" yaml:"title"`
	Closed     int64  `boil:"closed" json:"closed" toml:"closed" yaml:"closed"`
	Number     int64  `boil:"number" json:"number" toml:"number" yaml:"number"`
	Repository string `boil:"repository" json:"repository" toml:"repository" yaml:"repository"`

	R *issueR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L issueL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var IssueColumns = struct {
	ID         string
	URL        string
	Title      string
	Closed     string
	Number     string
	Repository string
}{
	ID:         "id",
	URL:        "url",
	Title:      "title",
	Closed:     "closed",
	Number:     "number",
	Repository: "repository",
}

var IssueTableColumns = struct {
	ID         string
	URL        string
	Title      string
	Closed     string
	Number     string
	Repository string
}{
	ID:         "issues.id",
	URL:        "issues.url",
	Title:      "issues.title",
	Closed:     "issues.closed",
	Number:     "issues.number",
	Repository: "issues.repository",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var IssueWhere = struct {
	ID         whereHelperstring
	URL        whereHelperstring
	Title      whereHelperstring
	Closed     whereHelperint64
	Number     whereHelperint64
	Repository whereHelperstring
}{
	ID:         whereHelperstring{field: "\"issues\".\"id\""},
	URL:        whereHelperstring{field: "\"issues\".\"url\""},
	Title:      whereHelperstring{field: "\"issues\".\"title\""},
	Closed:     whereHelperint64{field: "\"issues\".\"closed\""},
	Number:     whereHelperint64{field: "\"issues\".\"number\""},
	Repository: whereHelperstring{field: "\"issues\".\"repository\""},
}

// IssueRels is where relationship names are stored.
var IssueRels = struct {
	IssueRepository string
	Projectcards    string
}{
	IssueRepository: "IssueRepository",
	Projectcards:    "Projectcards",
}

// issueR is where relationships are stored.
type issueR struct {
	IssueRepository *Repository      `boil:"IssueRepository" json:"IssueRepository" toml:"IssueRepository" yaml:"IssueRepository"`
	Projectcards    ProjectcardSlice `boil:"Projectcards" json:"Projectcards" toml:"Projectcards" yaml:"Projectcards"`
}

// NewStruct creates a new relationship struct
func (*issueR) NewStruct() *issueR {
	return &issueR{}
}

func (r *issueR) GetIssueRepository() *Repository {
	if r == nil {
		return nil
	}
	return r.IssueRepository
}

func (r *issueR) GetProjectcards() ProjectcardSlice {
	if r == nil {
		return nil
	}
	return r.Projectcards
}

// issueL is where Load methods for each relationship are stored.
type issueL struct{}

var (
	issueAllColumns            = []string{"id", "url", "title", "closed", "number", "repository"}
	issueColumnsWithoutDefault = []string{"id", "url", "title", "number", "repository"}
	issueColumnsWithDefault    = []string{"closed"}
	issuePrimaryKeyColumns     = []string{"id"}
	issueGeneratedColumns      = []string{}
)

type (
	// IssueSlice is an alias for a slice of pointers to Issue.
	// This should almost always be used instead of []Issue.
	IssueSlice []*Issue
	// IssueHook is the signature for custom Issue hook methods
	IssueHook func(context.Context, boil.ContextExecutor, *Issue) error

	issueQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	issueType                 = reflect.TypeOf(&Issue{})
	issueMapping              = queries.MakeStructMapping(issueType)
	issuePrimaryKeyMapping, _ = queries.BindMapping(issueType, issueMapping, issuePrimaryKeyColumns)
	issueInsertCacheMut       sync.RWMutex
	issueInsertCache          = make(map[string]insertCache)
	issueUpdateCacheMut       sync.RWMutex
	issueUpdateCache          = make(map[string]updateCache)
	issueUpsertCacheMut       sync.RWMutex
	issueUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var issueAfterSelectHooks []IssueHook

var issueBeforeInsertHooks []IssueHook
var issueAfterInsertHooks []IssueHook

var issueBeforeUpdateHooks []IssueHook
var issueAfterUpdateHooks []IssueHook

var issueBeforeDeleteHooks []IssueHook
var issueAfterDeleteHooks []IssueHook

var issueBeforeUpsertHooks []IssueHook
var issueAfterUpsertHooks []IssueHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Issue) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Issue) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Issue) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Issue) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Issue) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Issue) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Issue) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Issue) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Issue) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range issueAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddIssueHook registers your hook function for all future operations.
func AddIssueHook(hookPoint boil.HookPoint, issueHook IssueHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		issueAfterSelectHooks = append(issueAfterSelectHooks, issueHook)
	case boil.BeforeInsertHook:
		issueBeforeInsertHooks = append(issueBeforeInsertHooks, issueHook)
	case boil.AfterInsertHook:
		issueAfterInsertHooks = append(issueAfterInsertHooks, issueHook)
	case boil.BeforeUpdateHook:
		issueBeforeUpdateHooks = append(issueBeforeUpdateHooks, issueHook)
	case boil.AfterUpdateHook:
		issueAfterUpdateHooks = append(issueAfterUpdateHooks, issueHook)
	case boil.BeforeDeleteHook:
		issueBeforeDeleteHooks = append(issueBeforeDeleteHooks, issueHook)
	case boil.AfterDeleteHook:
		issueAfterDeleteHooks = append(issueAfterDeleteHooks, issueHook)
	case boil.BeforeUpsertHook:
		issueBeforeUpsertHooks = append(issueBeforeUpsertHooks, issueHook)
	case boil.AfterUpsertHook:
		issueAfterUpsertHooks = append(issueAfterUpsertHooks, issueHook)
	}
}

// One returns a single issue record from the query.
func (q issueQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Issue, error) {
	o := &Issue{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for issues")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Issue records from the query.
func (q issueQuery) All(ctx context.Context, exec boil.ContextExecutor) (IssueSlice, error) {
	var o []*Issue

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to Issue slice")
	}

	if len(issueAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Issue records in the query.
func (q issueQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count issues rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q issueQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if issues exists")
	}

	return count > 0, nil
}

// IssueRepository pointed to by the foreign key.
func (o *Issue) IssueRepository(mods ...qm.QueryMod) repositoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.Repository),
	}

	queryMods = append(queryMods, mods...)

	return Repositories(queryMods...)
}

// Projectcards retrieves all the projectcard's Projectcards with an executor.
func (o *Issue) Projectcards(mods ...qm.QueryMod) projectcardQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"projectcards\".\"issue\"=?", o.ID),
	)

	return Projectcards(queryMods...)
}

// LoadIssueRepository allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (issueL) LoadIssueRepository(ctx context.Context, e boil.ContextExecutor, singular bool, maybeIssue interface{}, mods queries.Applicator) error {
	var slice []*Issue
	var object *Issue

	if singular {
		var ok bool
		object, ok = maybeIssue.(*Issue)
		if !ok {
			object = new(Issue)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeIssue)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeIssue))
			}
		}
	} else {
		s, ok := maybeIssue.(*[]*Issue)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeIssue)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeIssue))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &issueR{}
		}
		args = append(args, object.Repository)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &issueR{}
			}

			for _, a := range args {
				if a == obj.Repository {
					continue Outer
				}
			}

			args = append(args, obj.Repository)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`repositories`),
		qm.WhereIn(`repositories.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Repository")
	}

	var resultSlice []*Repository
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Repository")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for repositories")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for repositories")
	}

	if len(repositoryAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.IssueRepository = foreign
		if foreign.R == nil {
			foreign.R = &repositoryR{}
		}
		foreign.R.Issues = append(foreign.R.Issues, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.Repository == foreign.ID {
				local.R.IssueRepository = foreign
				if foreign.R == nil {
					foreign.R = &repositoryR{}
				}
				foreign.R.Issues = append(foreign.R.Issues, local)
				break
			}
		}
	}

	return nil
}

// LoadProjectcards allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (issueL) LoadProjectcards(ctx context.Context, e boil.ContextExecutor, singular bool, maybeIssue interface{}, mods queries.Applicator) error {
	var slice []*Issue
	var object *Issue

	if singular {
		var ok bool
		object, ok = maybeIssue.(*Issue)
		if !ok {
			object = new(Issue)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeIssue)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeIssue))
			}
		}
	} else {
		s, ok := maybeIssue.(*[]*Issue)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeIssue)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeIssue))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &issueR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &issueR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`projectcards`),
		qm.WhereIn(`projectcards.issue in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load projectcards")
	}

	var resultSlice []*Projectcard
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice projectcards")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on projectcards")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for projectcards")
	}

	if len(projectcardAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Projectcards = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &projectcardR{}
			}
			foreign.R.ProjectcardIssue = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.Issue) {
				local.R.Projectcards = append(local.R.Projectcards, foreign)
				if foreign.R == nil {
					foreign.R = &projectcardR{}
				}
				foreign.R.ProjectcardIssue = local
				break
			}
		}
	}

	return nil
}

// SetIssueRepository of the issue to the related item.
// Sets o.R.IssueRepository to related.
// Adds o to related.R.Issues.
func (o *Issue) SetIssueRepository(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Repository) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"issues\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"repository"}),
		strmangle.WhereClause("\"", "\"", 0, issuePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.Repository = related.ID
	if o.R == nil {
		o.R = &issueR{
			IssueRepository: related,
		}
	} else {
		o.R.IssueRepository = related
	}

	if related.R == nil {
		related.R = &repositoryR{
			Issues: IssueSlice{o},
		}
	} else {
		related.R.Issues = append(related.R.Issues, o)
	}

	return nil
}

// AddProjectcards adds the given related objects to the existing relationships
// of the issue, optionally inserting them as new records.
// Appends related to o.R.Projectcards.
// Sets related.R.ProjectcardIssue appropriately.
func (o *Issue) AddProjectcards(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Projectcard) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.Issue, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"projectcards\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 0, []string{"issue"}),
				strmangle.WhereClause("\"", "\"", 0, projectcardPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.Issue, o.ID)
		}
	}

	if o.R == nil {
		o.R = &issueR{
			Projectcards: related,
		}
	} else {
		o.R.Projectcards = append(o.R.Projectcards, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &projectcardR{
				ProjectcardIssue: o,
			}
		} else {
			rel.R.ProjectcardIssue = o
		}
	}
	return nil
}

// SetProjectcards removes all previously related items of the
// issue replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.ProjectcardIssue's Projectcards accordingly.
// Replaces o.R.Projectcards with related.
// Sets related.R.ProjectcardIssue's Projectcards accordingly.
func (o *Issue) SetProjectcards(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Projectcard) error {
	query := "update \"projectcards\" set \"issue\" = null where \"issue\" = ?"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.Projectcards {
			queries.SetScanner(&rel.Issue, nil)
			if rel.R == nil {
				continue
			}

			rel.R.ProjectcardIssue = nil
		}
		o.R.Projectcards = nil
	}

	return o.AddProjectcards(ctx, exec, insert, related...)
}

// RemoveProjectcards relationships from objects passed in.
// Removes related items from R.Projectcards (uses pointer comparison, removal does not keep order)
// Sets related.R.ProjectcardIssue.
func (o *Issue) RemoveProjectcards(ctx context.Context, exec boil.ContextExecutor, related ...*Projectcard) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.Issue, nil)
		if rel.R != nil {
			rel.R.ProjectcardIssue = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("issue")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.Projectcards {
			if rel != ri {
				continue
			}

			ln := len(o.R.Projectcards)
			if ln > 1 && i < ln-1 {
				o.R.Projectcards[i] = o.R.Projectcards[ln-1]
			}
			o.R.Projectcards = o.R.Projectcards[:ln-1]
			break
		}
	}

	return nil
}

// Issues retrieves all the records using an executor.
func Issues(mods ...qm.QueryMod) issueQuery {
	mods = append(mods, qm.From("\"issues\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"issues\".*"})
	}

	return issueQuery{q}
}

// FindIssue retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindIssue(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Issue, error) {
	issueObj := &Issue{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"issues\" where \"id\"=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, issueObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from issues")
	}

	if err = issueObj.doAfterSelectHooks(ctx, exec); err != nil {
		return issueObj, err
	}

	return issueObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Issue) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no issues provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(issueColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	issueInsertCacheMut.RLock()
	cache, cached := issueInsertCache[key]
	issueInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			issueAllColumns,
			issueColumnsWithDefault,
			issueColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(issueType, issueMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(issueType, issueMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"issues\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"issues\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into issues")
	}

	if !cached {
		issueInsertCacheMut.Lock()
		issueInsertCache[key] = cache
		issueInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Issue.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Issue) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	issueUpdateCacheMut.RLock()
	cache, cached := issueUpdateCache[key]
	issueUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			issueAllColumns,
			issuePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update issues, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"issues\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, issuePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(issueType, issueMapping, append(wl, issuePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update issues row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for issues")
	}

	if !cached {
		issueUpdateCacheMut.Lock()
		issueUpdateCache[key] = cache
		issueUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q issueQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for issues")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for issues")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o IssueSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), issuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"issues\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, issuePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in issue slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all issue")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Issue) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no issues provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(issueColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	issueUpsertCacheMut.RLock()
	cache, cached := issueUpsertCache[key]
	issueUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			issueAllColumns,
			issueColumnsWithDefault,
			issueColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			issueAllColumns,
			issuePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("db: unable to upsert issues, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(issuePrimaryKeyColumns))
			copy(conflict, issuePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"issues\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(issueType, issueMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(issueType, issueMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "db: unable to upsert issues")
	}

	if !cached {
		issueUpsertCacheMut.Lock()
		issueUpsertCache[key] = cache
		issueUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Issue record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Issue) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no Issue provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), issuePrimaryKeyMapping)
	sql := "DELETE FROM \"issues\" WHERE \"id\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from issues")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for issues")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q issueQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no issueQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from issues")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for issues")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o IssueSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(issueBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), issuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"issues\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, issuePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from issue slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for issues")
	}

	if len(issueAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Issue) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindIssue(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *IssueSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := IssueSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), issuePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"issues\".* FROM \"issues\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, issuePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in IssueSlice")
	}

	*o = slice

	return nil
}

// IssueExists checks if the Issue row exists.
func IssueExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"issues\" where \"id\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if issues exists")
	}

	return exists, nil
}

// Exists checks if the Issue row exists.
func (o *Issue) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return IssueExists(ctx, exec, o.ID)
}
