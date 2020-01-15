<template>
  <div class="container">
    <common-head></common-head>
    <div class="ah-carousel">
      <el-carousel height="400px">
        <el-carousel-item>
          <img src="./imgs/banner/xm.png" alt="">
        </el-carousel-item>
        <el-carousel-item>
          <img src="./imgs/banner/ym.png" alt="">
        </el-carousel-item>
      </el-carousel>
    </div>
    <router-view class="common-content"></router-view>
  </div>
</template>
<script>
import CommonHead from "@/components/header";
import CommonFooter from "@/components/footer";
export default {
  name: "app",
  components: {
    CommonHead,
    CommonFooter
  }
};
</script>
<style lang="less" scoped>
.container {
  width: 100%;
  .common-content{
    margin: 0 auto;
    width: 70%;
  }
  .ah-carousel {
    margin: 20px auto;
    .el-carousel__item {

    }
    img {
      margin: 0 auto;
      width: 100%;
      height: 400px;
    }
  }
}
</style>

