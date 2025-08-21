package repository

import "errors"

var (
    ErrFailedToFetch     = errors.New("Failed to fetch data")
    ErrFailedToInsert    = errors.New("Failed to insert data")
    ErrFailedToDelete    = errors.New("Failed to delete data")
    ErrFailedToScan      = errors.New("Failed to scan data")
    ErrAdminNotFound     = errors.New("Admin not found")
    ErrCreateAdmin       = errors.New("Failed to create admin")
    ErrIteration         = errors.New("Error iterating")
    ErrActivityNotFound  = errors.New("Activity not found")
    ErrFetchActivity     = errors.New("Failed to fetch activity")
    ErrAllActivities     = errors.New("Failed to fetch all activities")
    ErrCreateActivity    = errors.New("Failed to create activity")
    ErrSaveActivity      = errors.New("Could not save activity")
    ErrLinkActivity      = errors.New("Could not link prompt to activity")
    ErrRemoveLink        = errors.New("Could not remove link")
    ErrUpdateActivity    = errors.New("Failed to update activity")
    ErrSaveSession       = errors.New("Could not save session")
    ErrDeleteSession     = errors.New("Can't delete session")
    ErrDuplicateActivity = errors.New("This activity already exists")
)
