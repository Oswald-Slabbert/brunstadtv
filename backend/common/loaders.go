package common

import (
	"github.com/bcc-code/brunstadtv/backend/batchloaders"
	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader/v7"
)

// BatchLoaders contains loaders for the different items
type BatchLoaders struct {
	ApplicationLoader                  *dataloader.Loader[int, *Application]
	ApplicationIDFromCodeLoader        *dataloader.Loader[string, *int]
	RedirectLoader                     *dataloader.Loader[uuid.UUID, *Redirect]
	RedirectIDFromCodeLoader           *dataloader.Loader[string, *uuid.UUID]
	PageLoader                         *dataloader.Loader[int, *Page]
	PageIDFromCodeLoader               *dataloader.Loader[string, *int]
	CollectionIDFromSlugLoader         *batchloaders.BatchLoader[string, *int]
	SectionLoader                      *dataloader.Loader[int, *Section]
	SectionsLoader                     *dataloader.Loader[int, []*int]
	CollectionLoader                   *dataloader.Loader[int, *Collection]
	CollectionItemLoader               *dataloader.Loader[int, []*CollectionItem]
	StudyTopicLoader                   *batchloaders.BatchLoader[uuid.UUID, *StudyTopic]
	StudyLessonLoader                  *batchloaders.BatchLoader[uuid.UUID, *Lesson]
	StudyTaskLoader                    *batchloaders.BatchLoader[uuid.UUID, *Task]
	StudyQuestionAlternativesLoader    *batchloaders.BatchLoader[uuid.UUID, []*QuestionAlternative]
	ShowLoader                         *dataloader.Loader[int, *Show]
	SeasonLoader                       *dataloader.Loader[int, *Season]
	EpisodeLoader                      *dataloader.Loader[int, *Episode]
	EpisodeIDFromLegacyIDLoader        *dataloader.Loader[int, *int]
	EpisodeIDFromLegacyProgramIDLoader *dataloader.Loader[int, *int]
	LinkLoader                         *dataloader.Loader[int, *Link]
	FilesLoader                        *dataloader.Loader[int, []*File]
	StreamsLoader                      *dataloader.Loader[int, []*Stream]
	EventLoader                        *batchloaders.BatchLoader[int, *Event]
	FAQCategoryLoader                  *dataloader.Loader[int, *FAQCategory]
	QuestionLoader                     *dataloader.Loader[int, *Question]
	QuestionsLoader                    *dataloader.Loader[int, []*int]
	ProfilesLoader                     *dataloader.Loader[string, []*Profile]
	MessageGroupLoader                 *dataloader.Loader[int, *MessageGroup]
	RedirectFromCodeLoader             *dataloader.Loader[string, *Redirect]

	EpisodeProgressLoader *batchloaders.BatchLoader[uuid.UUID, []*int]
	// Permissions
	ShowPermissionLoader    *dataloader.Loader[int, *Permissions[int]]
	SeasonPermissionLoader  *dataloader.Loader[int, *Permissions[int]]
	EpisodePermissionLoader *dataloader.Loader[int, *Permissions[int]]
	PagePermissionLoader    *dataloader.Loader[int, *Permissions[int]]
	SectionPermissionLoader *dataloader.Loader[int, *Permissions[int]]

	CompletedTopicsLoader         *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]
	CompletedLessonsLoader        *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]
	CompletedTasksLoader          *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]
	CompletedAndLockedTasksLoader *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]

	// Achievements
	AchievementLoader             *batchloaders.BatchLoader[uuid.UUID, *Achievement]
	AchievementGroupLoader        *batchloaders.BatchLoader[uuid.UUID, *AchievementGroup]
	AchievementsLoader            *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]
	UnconfirmedAchievementsLoader *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]

	AchievementGroupAchievementsLoader *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]

	ComputedDataLoader *batchloaders.BatchLoader[uuid.UUID, []*ComputedData]
}

// FilteredLoaders contains loaders that will be filtered by permissions.
type FilteredLoaders struct {
	EpisodeFilterLoader     *dataloader.Loader[int, *int]
	EpisodesLoader          *dataloader.Loader[int, []*int]
	SeasonFilterLoader      *dataloader.Loader[int, *int]
	SeasonsLoader           *dataloader.Loader[int, []*int]
	ShowFilterLoader        *dataloader.Loader[int, *int]
	SectionsLoader          *dataloader.Loader[int, []*int]
	CollectionItemsLoader   *dataloader.Loader[int, []*CollectionItem]
	CollectionItemIDsLoader *dataloader.Loader[int, []Identifier]
	CalendarEntryLoader     *batchloaders.BatchLoader[int, *CalendarEntry]
	StudyTopicFilterLoader  *batchloaders.BatchLoader[uuid.UUID, *uuid.UUID]
	StudyLessonsLoader      *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]
	StudyLessonFilterLoader *batchloaders.BatchLoader[uuid.UUID, *uuid.UUID]
	StudyTasksLoader        *batchloaders.BatchLoader[uuid.UUID, []*uuid.UUID]
	StudyTaskFilterLoader   *batchloaders.BatchLoader[uuid.UUID, *uuid.UUID]

	//Relations
	StudyLessonEpisodesLoader *batchloaders.BatchLoader[uuid.UUID, []*int]
	EpisodeStudyLessonsLoader *batchloaders.BatchLoader[int, []*uuid.UUID]
	StudyLessonLinksLoader    *batchloaders.BatchLoader[uuid.UUID, []*int]
	LinkStudyLessonsLoader    *batchloaders.BatchLoader[int, []*uuid.UUID]
}

// ProfileLoaders contains loaders per profile
type ProfileLoaders struct {
	ProgressLoader                *batchloaders.BatchLoader[int, *Progress]
	TaskCompletedLoader           *batchloaders.BatchLoader[uuid.UUID, *uuid.UUID]
	AchievementAchievedAtLoader   *batchloaders.BatchLoader[uuid.UUID, *Achieved]
	GetSelectedAlternativesLoader *batchloaders.BatchLoader[uuid.UUID, *SelectedAlternatives]
}
