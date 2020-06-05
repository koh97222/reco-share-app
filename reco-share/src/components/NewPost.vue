<template>
  <div class="container">
    <b-modal centered size="lg" ref="new-post" ok-only>
      <h2 class="text-center">NEW POST</h2>
      <b-alert variant="danger" show dismissible fade v-if="isError"
        >投稿に失敗しました。</b-alert
      >
      <b-input-group prepend="title" class="mt-4">
        <b-input placeholder="タイトル" v-model="post.title"></b-input>
      </b-input-group>
      <b-input-group class="mt-1" prepend="tag">
        <b-input placeholder="タグ" v-model="post.tag"></b-input>
      </b-input-group>
      <b-input-group class="mt-1" prepend="message">
        <b-form-textarea
          placeholder="出会った感想などご自由に♪"
          v-model="post.description"
        ></b-form-textarea>
      </b-input-group>
      <template v-slot:modal-footer>
        <b-button
          size="sm"
          variant="secondary"
          @click="submit(post)"
          :disabled="disabledButton"
        >
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
      this.init(this.post);
      this.$refs["new-post"].show();
    });
    // TODO:cookieからuserIDを取得し、格納する。
    let el = document.cookie.split("=");
    this.post.UserID = Number(el[1]);
  },
  data() {
    return {
      post: {
        UserID: null,
        title: null,
        tag: null,
        description: null,
      },
      isError: false,
    };
  },
  computed: {
    disabledButton() {
      return this.post.title == null || this.post.title == "";
    },
  },
  methods: {
    init(post) {
      // スマートに書きたい
      post.UserID = null;
      post.title = null;
      post.tag = null;
      post.description = null;
    },
    submit(post) {
      let _self = this;
      _self.$axios.post("http://localhost:8082/postarticle", post).then((r) => {
        if (r.status == 200) {
          if (r.data.resultCode == "80") {
            this.isError = true;
          } else {
            console.log("登録成功！！");
          }
        }
      });
    },
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
