<template>
  <div class="container">
    <b-row>
      <b-col xl="2"></b-col>
      <b-col xl="8" cols="12">
        <b-card class="post">
          <b-row>
            <b-col xl="2" cols="3"><b-badge>title</b-badge></b-col>
            <b-col xl="10" cols="9"
              ><h6 class="text-left">
                <b-link class="text-left">最近買ったターンテーブル</b-link>
              </h6></b-col
            >
          </b-row>
          <b-row>
            <b-col xl="2" cols="3"><b-badge>tag</b-badge></b-col>
            <b-col xl="10" cols="9"
              ><h6 class="text-left">
                <b-link class="text-left">#record </b-link>
                <b-link class="text-left">#vinyl </b-link>
                <b-link class="text-left">#tokyo </b-link>
              </h6></b-col
            >
          </b-row>
          <b-row>
            <b-col xl="2" cols="3"><b-badge>user</b-badge></b-col>
            <b-col xl="10" cols="9"
              ><h6 class="text-left">
                <b-link class="text-left">@user</b-link>
              </h6></b-col
            >
          </b-row>
          <hr />
          <div class="img">
            image
          </div>
          <template v-slot:footer class="message mt-4">
            <p>textが入ります。textが入ります。textが入ります。</p>
          </template>
        </b-card>
      </b-col>
      <b-col xl="2"></b-col>
    </b-row>
  </div>
</template>
<script>
export default {};
</script>
<style scoped>
.post {
  box-shadow: 2px 2px 4px #cccccc;
  transition: 0.3s;
}
.post:hover {
  transform: scale(1.03, 1.03);
}
.card-footer {
  color: #cccccc;
  background-color: #10100e;
  box-shadow: 2px 2px 4px #cccccc;
  opacity: 0.9;
}
</style>
