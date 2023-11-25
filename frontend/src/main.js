import Vue from "vue"
import App from "./App.vue"
import "./registerServiceWorker"
import router from "./router"
import store from "./store"
import vuetify from "./plugins/vuetify"
import posthogPlugin from "./plugins/posthog"
import VueGtm from "@gtm-support/vue2-gtm"
import VueMeta from "vue-meta"
import "./index.css"

// Posthog
if (process.env.NODE_ENV !== "development") {
  Vue.use(posthogPlugin)
}

// Google Analytics
Vue.use(VueGtm, {
  id: "GTM-M677X6V",
  vueRouter: router,
})

// Site Metadata
Vue.use(VueMeta)

Vue.config.productionTip = false

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount("#app")
