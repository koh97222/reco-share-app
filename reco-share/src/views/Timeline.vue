<template>
  <div class="timeline">
    <div class="container">
      <div class="title">
        <h2 class="fw-700">FEATURED</h2>
        <hr />
        <post class="mt-5"></post>
      </div>
      <div class="latest-post title">
        <h2 class="fw-700">LATEST POST</h2>
        <hr />
        <post class="mt-5"></post>
        <post class="mt-5"></post>
        <post class="mt-5"></post>
        <b-button
          class="mt-5 mb-5"
          style="background-color:black; font-weight:400"
          >VIEW MORE</b-button
        >
      </div>
    </div>
  </div>
</template>

<script>
import Post from "../components/Post.vue";
export default {
  components: { Post },
};
</script>

<style>
@media only screen and (max-width: 750px) {
  img {
    max-width: 100%;
  }
}
.badge-secondary {
  width: 60px;
}
.btn {
  width: 300px;
  height: 50px;
}
.fw-700 {
  font-weight: 700;
}
.timeline {
  background-color: #f2f2f2;
}
.title {
  padding: 15px;
}
</style>
