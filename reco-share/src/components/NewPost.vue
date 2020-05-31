<template>
  <div class="container">
    <b-modal centered size="lg" ref="new-post" ok-only>
      <h2 class="text-center">NEW POST</h2>
      <b-input-group prepend="title" class="mt-4">
        <b-input placeholder="タイトル"></b-input>
      </b-input-group>
      <b-input-group class="mt-1" prepend="tag">
        <b-input placeholder="タグ"></b-input>
      </b-input-group>
      <b-input-group class="mt-1" prepend="message">
        <b-form-textarea
          placeholder="出会った感想などご自由に♪"
        ></b-form-textarea>
      </b-input-group>
      <template v-slot:modal-footer>
        <b-button size="sm" variant="secondary">
          投稿する <font-awesome-icon icon="blog" />
        </b-button>
      </template>
    </b-modal>
  </div>
</template>
<script>
import EventBus from "../eventbus.js";
export default {
  created() {
    EventBus.$on("open-newpost-modal", () => {
      this.$refs["new-post"].show();
    });
  },
};
</script>
<style scoped>
.btn {
  color: black;
  background-color: white;
  box-shadow: 2px 2px 4px #cccccc;
  width: 170px;
}
.btn:hover {
  background-color: #e5b848;
  color: black;
}
.input-group-text {
  width: 100px;
}
</style>
