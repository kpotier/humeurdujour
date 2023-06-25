import { visible } from "@/stores/progressBar";
import {
  createRouter,
  createWebHistory,
  type RouteLocationNormalized,
} from "vue-router";
import HomeView from "../views/HomeView.vue";
import Error404View from "../views/Error404View.vue";
import settings from "./settings";
import DashboardVueVue from "@/views/DashboardVue.vue";
import GraphesVueVue from "@/views/GraphesVue.vue";

// Note: by default, requiresAuth is true.

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
      redirect: { name: "dashboard" },
      children: [
        settings,
        {
          name: "dashboard",
          path: "dashboard",
          component: DashboardVueVue,
        },
        {
          name: "graphes",
          path: "graphes",
          component: GraphesVueVue,
        },
      ],
    },
    {
      path: "/:pathMatch(.*)*",
      component: Error404View,
    },
  ],
});

router.beforeEach(async (to) => {
  visible.value = { to: to, count: visible.value.count + 1 };
});

router.afterEach((to, from, failed) => {
  if (
    failed?.type !== 8 &&
    (visible.value.to === to || visible.value.to == to.redirectedFrom)
  )
    visible.value = { to: {} as RouteLocationNormalized, count: 0 };
});

export default router;
