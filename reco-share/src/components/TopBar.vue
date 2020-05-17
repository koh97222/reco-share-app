<template>
  <div class="top-bar">
    <b-navbar class="gothic" toggleable="lg" style="background-color:white;">
      <b-navbar-brand tag="h1" href="#">
        <div class="ml-5">
          reco<span class="e5b848">×</span>share
          <font-awesome-icon icon="headphones-alt" class="ml-1" />
        </div>
      </b-navbar-brand>
      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>
      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav class="ml-auto">
          <div class="mr-4">
            <b-nav-item href="#">
              Sign In
              <font-awesome-icon icon="sign-in-alt"
            /></b-nav-item>
          </div>
          <div class="mr-4">
            <b-nav-item href="#" @click="showRegistModal()">
              新規ユーザ登録<font-awesome-icon icon="user-plus" class="ml-1" />
            </b-nav-item>
          </div>
        </b-navbar-nav>
      </b-collapse>
    </b-navbar>
    <user-registration ref="registModal"></user-registration>
  </div>
</template>
<script>
import UserRegistration from "./UserRegistration.vue";

export default {
  components: { UserRegistration },
  methods: {
    showRegistModal() {
      this.$refs.registModal.$refs["user-regist"].show();
    },
  },
};
</script>
<style scoped>
.e5b848 {
  color: #e5b848;
}
.gothic {
  font-family: "游ゴシック Medium", "Yu Gothic Medium", "游ゴシック体",
    "YuGothic", "ヒラギノ角ゴ ProN W3", "Hiragino Kaku Gothic ProN", "メイリオ",
    "Meiryo", "verdana", sans-serif;
}
.items:hover {
  background-color: #e5b848;
  box-shadow: 2px 2px 4px #cccccc;
}
.nav-link {
  -webkit-transition: all 0.3s ease;
  -moz-transition: all 0.3s ease;
  -o-transition: all 0.3s ease;
  transition: all 0.3s ease;
}
.nav-link:hover {
  background-color: #e5b848;
  box-shadow: 2px 2px 4px #cccccc;
}
.navbar-brand {
  font-weight: 700;
}
.navbar-expand-lg {
  border-bottom: 1px solid;
}
.ml-5 {
  padding: 5px;
}
</style>
