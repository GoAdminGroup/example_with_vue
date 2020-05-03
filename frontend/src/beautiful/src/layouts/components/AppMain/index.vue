<template>
  <section class="app-main-container">
    <transition mode="out-in" name="fade-transform">
      <keep-alive :include="cachedViews" :max="10">
        <byui-keel v-if="show" style="margin: 15px;">
          <byui-keel-heading :img="true" />
          <byui-keel-text :lines="7" />
          <byui-keel-heading :img="true" />
          <byui-keel-text :lines="6" />
          <byui-keel-heading :img="true" />
          <byui-keel-text :lines="8" />
        </byui-keel>
        <router-view :key="key" style="min-height: 78vh;" />
      </keep-alive>
    </transition>
  </section>
</template>

<script>
import { ByuiKeel, ByuiKeelHeading, ByuiKeelText } from "@/plugins/byuiKeel";
import { mapGetters } from "vuex";

export default {
  name: "AppMain",
  components: {
    ByuiKeel,
    ByuiKeelHeading,
    ByuiKeelText,
  },
  data() {
    return {
      show: true,
      fullYear: new Date().getFullYear(),
    };
  },
  computed: {
    ...mapGetters(["cachedViews", "device"]),
    key() {
      return this.$route.path;
    },
  },
  watch: {
    $route(to, from) {
      this.$nextTick(() => {
        /*if (this.$store.state.tagsView.skeleton) {
                          this.show = true;
                          setTimeout(() => {
                            this.show = false;
                          }, 0);
                        } else {
                          this.show = false;
                        }*/
        if ("mobile" === this.device) {
          this.$store.dispatch("settings/foldSideBar");
          $("body").attr("style", "");
        }
      });
    },
  },
  created() {},
  mounted() {
    this.$nextTick(() => {
      this.show = true;
      setTimeout(() => {
        this.show = false;
      }, 400);
    });
  },
  methods: {},
};
</script>

<style lang="scss" scoped>
.app-main-container {
  width: 100%;
  position: relative;
  overflow: hidden;

  .footer-copyright {
    font-size: $base-font-size-small;
    height: 70px;
    line-height: 30px;
    text-align: center;
    color: rgba(0, 0, 0, 0.45);
  }
}
</style>
