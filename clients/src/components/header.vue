<template>
  <div class="common-header">
    <div class="header-left">
      <a class="title-text" href="#/home">
        <span>昂和水溶肥</span>
      </a>
    </div>
    <div class="header-right">
      <div class="nav-list">
        <router-link class="nav-text" to="/home">
          <span>首页</span>
        </router-link>
        <router-link class="nav-text" to="/about">
          <span>关于我们</span>
        </router-link>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  mounted() {
    console.log(this.$route);
  }
};
</script>
<style lang="less">
.common-header {
  width: 70%;
  height: 60px;
  margin: 0 auto;
  background: #fff;
  display: flex;
  justify-content: space-between;
  .router-link-active {
    color: #f85415!important;
    border-bottom: 2px solid #f85415;
  }
  .header-left {
    .title-text {
      color: #409EFF;
      line-height: 60px;
      font-size: 26px;
    }
  }
  .header-right {
    .nav-list {
      display: flex;
      justify-content: space-between;
      .nav-text {
        color: #409EFF;
        line-height: 60px;
        font-size: 16px;
        margin-left: 20px;
      }
    }
  }
}
</style>
