<template>
    <section
        class="max-w-screen-2xl mx-auto rounded rounded-2xl"
        v-if="episode"
    >
        <div class="relative aspect-video w-full">
            <div
                class="h-full w-full bg-secondary rounded rounded-xl opacity-10 absolute"
            ></div>
            <EpisodeViewer
                :context="context"
                :auto-play="autoPlay"
                class="drop-shadow-xl overflow-hidden"
                :episode="episode"
            ></EpisodeViewer>
        </div>
        <div class="flex flex-col">
            <div class="bg-primary-light p-4 w-full">
                <div class="flex">
                    <h1 class="text-lg lg:text-xl font-medium">
                        {{ episode.title }}
                    </h1>
                    <div class="ml-auto">
                        <SharePopover :episode="episode"></SharePopover>
                    </div>
                </div>
                <div class="flex">
                    <h1 class="my-auto flex gap-1">
                        <AgeRating
                            :episode="episode"
                            :show-a="true"
                            class="mr-1"
                        />
                        <span v-if="episode.season" class="text-primary">{{
                            episode.season.show.title
                        }}</span>
                        <span
                            v-if="episode.productionDateInTitle"
                            class="text-gray ml-1"
                            >{{
                                new Date(episode.productionDate).toDateString()
                            }}</span
                        >
                    </h1>
                </div>
                <div class="text-white mt-2 opacity-70 text-md lg:text-lg">
                    {{ episode.description }}
                </div>
                <LessonButton
                    v-if="lesson && !episodeComingSoon(episode)"
                    class="mt-4"
                    :lesson="lesson"
                    :episode-id="episode.id"
                    @click="openLesson"
                />
            </div>
            <div class="flex flex-col gap-2 mt-4">
                <div class="flex gap-4 p-2 font-semibold">
                    <button
                        v-if="
                            episode.context?.__typename === 'ContextCollection'
                        "
                        class="bg-primary-light uppercase border-gray border px-3 py-1 rounded-full transition duration-100"
                        :class="[
                            effectiveView === 'context'
                                ? 'opacity-100 border-opacity-40 '
                                : 'opacity-50 bg-opacity-0 border-opacity-0',
                        ]"
                        @click="effectiveView = 'context'"
                    >
                        {{ $t("episode.videos") }}
                    </button>
                    <button
                        v-else-if="seasonId"
                        class="bg-primary-light uppercase border-gray border px-3 py-1 rounded-full transition duration-100"
                        :class="[
                            effectiveView === 'episodes'
                                ? 'opacity-100 border-opacity-40 '
                                : 'opacity-50 bg-opacity-0 border-opacity-0',
                        ]"
                        @click="effectiveView = 'episodes'"
                    >
                        {{ $t("episode.episodes") }}
                    </button>
                    <button
                        class="bg-primary-light uppercase border-gray border px-3 py-1 rounded-full transition duration-100"
                        :class="[
                            effectiveView === 'details'
                                ? 'opacity-100 border-opacity-40'
                                : 'opacity-50 bg-opacity-0 border-opacity-0',
                        ]"
                        @click="effectiveView = 'details'"
                    >
                        {{ $t("episode.details") }}
                    </button>
                </div>
                <hr class="border-gray border-opacity-70" />
                <div>
                    <Transition name="slide-fade" mode="out-in">
                        <EpisodeDetails
                            v-if="effectiveView === 'details'"
                            :episode="episode"
                        ></EpisodeDetails>
                        <div v-else-if="effectiveView === 'context'">
                            <ItemList
                                :items="
                                    toListItems(
                                        episode.context?.__typename ===
                                            'ContextCollection'
                                            ? episode.context.items?.items ?? []
                                            : []
                                    )
                                "
                                :current-id="episode.id"
                                @item-click="(i) => setEpisode(i.id)"
                            ></ItemList>
                        </div>
                        <div
                            v-else-if="effectiveView === 'episodes'"
                            class="flex flex-col"
                        >
                            <SeasonSelector
                                v-if="episode.season"
                                :items="episode.season?.show.seasons.items"
                                v-model="seasonId"
                            ></SeasonSelector>
                            <ItemList
                                :items="seasonEpisodes"
                                :current-id="episode.id"
                                @item-click="(i) => setEpisode(i.id)"
                            ></ItemList>
                        </div>
                    </Transition>
                </div>
            </div>
        </div>
        <div v-if="error" class="text-red">{{ error.message }}</div>
    </section>
    <LoginToView v-else-if="noAccess && !authenticated"> </LoginToView>
    <NotFound v-else-if="!loading" :title="$t('episode.notFound')"></NotFound>
</template>
<script lang="ts" setup>
import {
    EpisodeContext,
    GetEpisodeQuery,
    GetSeasonOnEpisodePageQuery,
    useGetEpisodeQuery,
    useGetSeasonOnEpisodePageQuery,
} from "@/graph/generated"
import { computed, nextTick, ref, watch } from "vue"
import EpisodeViewer from "@/components/EpisodeViewer.vue"
import EpisodeDetails from "@/components/episodes/EpisodeDetails.vue"
import AgeRating from "@/components/episodes/AgeRating.vue"
import SeasonSelector from "@/components/SeasonSelector.vue"
import ItemList from "../sections/ItemList.vue"
import NotFound from "../NotFound.vue"
import LoginToView from "./LoginToView.vue"
import { episodesToListItems, toListItems } from "@/utils/lists"
import { useAuth } from "@/services/auth"
import SharePopover from "./SharePopover.vue"
import LessonButton from "../study/LessonButton.vue"
import router from "@/router"
import { episodeComingSoon } from "../../utils/items"

const props = defineProps<{
    initialEpisodeId: string
    context: EpisodeContext
    autoPlay?: boolean
}>()

const { authenticated } = useAuth()

const emit = defineEmits<{
    (e: "episode", v: GetEpisodeQuery["episode"]): void
}>()

const episode = ref(null as GetEpisodeQuery["episode"] | null)
const season = ref(null as GetSeasonOnEpisodePageQuery["season"] | null)

const seasonId = ref("")
const loading = ref(true)

const context = ref(props.context)

const episodeId = ref(props.initialEpisodeId)

const setEpisode = (id: string) => {
    episodeId.value = id
    nextTick()
        .then(load)
        .then(() => {
            if (episode.value) {
                emit("episode", episode.value)
            }
        })
}

const { error, executeQuery } = useGetEpisodeQuery({
    pause: true,
    variables: {
        episodeId,
        context,
    },
})

const lesson = computed(() => episode.value?.lessons.items[0])
const openLesson = () =>
    router.push("/ep/" + episode.value?.id + "/lesson/" + lesson.value?.id)

const noAccess = computed(() => {
    return error.value?.graphQLErrors.some(
        (e) => e.extensions.code === "item/no-access"
    )
})

const seasonQuery = useGetSeasonOnEpisodePageQuery({
    pause: true,
    variables: {
        seasonId,
        firstEpisodes: 50,
    },
})

const seasonEpisodes = computed(() => {
    return episodesToListItems(
        seasonQuery.data.value?.season.episodes.items ?? []
    )
})

const loadSeason = async () => {
    season.value = null
    if (seasonId.value) {
        const r = await seasonQuery.executeQuery()
        season.value = r.data.value?.season ?? null
    }
}

watch(() => seasonId.value, loadSeason)

const load = async () => {
    loading.value = true
    season.value = null
    seasonId.value = ""
    const r = await executeQuery()
    if (r.data.value?.episode) {
        episode.value = r.data.value.episode

        if (!context.value?.collectionId) {
            if (episode.value.season?.id) {
                seasonId.value = episode.value.season.id
                await nextTick()
            }
        }
        await loadSeason()
    }
    loading.value = false
}

load()
const view = ref(null as "episodes" | "details" | "context" | null)

const effectiveView = computed({
    get() {
        const v = view.value
        switch (v) {
            case "context":
                if (
                    episode.value?.context?.__typename === "ContextCollection"
                ) {
                    return "context"
                }
                break
            case "episodes":
                if (episode.value?.season) {
                    return "episodes"
                }
                break
            case "details":
                return "details"
        }

        if (episode.value?.context?.__typename === "ContextCollection") {
            return "context"
        }
        return (view.value = !episode.value?.season ? "details" : "episodes")
    },
    set(v) {
        view.value = v
    },
})
</script>
