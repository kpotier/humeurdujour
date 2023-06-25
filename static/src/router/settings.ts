export default {
  name: "settings",
  path: "settings",
  component: () => import("@/views/settings/LayoutVue.vue"),
  redirect: { name: "settingsAppearance" },
  children: [
    {
      name: "settingsAppearance",
      path: "appearance",
      component: () => import("@/views/settings/AppearanceVue.vue"),
    },
  ],
};
