import { createRouter, createWebHistory } from "vue-router";
import Home from "../views/Home.vue";
import ModuleDetails from "../views/ModuleDetails.vue";
import NotFound from "../views/NotFound.vue";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/details/:code",
    name: "ModuleDetails",
    component: ModuleDetails,
    props: true,
  },
  //catchall 404
  {
    path: "/:catchAll(.*)",
    name: "NotFound",
    component: NotFound,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
