import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import "bootstrap/dist/css/bootstrap.min.css";
import TabNav from "./components/TabNav.vue";
import Tab from "./components/Tab.vue";

createApp(App)
  .use(router)
  .component("tab-nav", TabNav)
  .component("tab", Tab)
  .mount("#app");
