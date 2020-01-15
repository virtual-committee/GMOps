<template>
  <div class="home-carousel">
    <el-carousel type="card" height="400px">
      <el-carousel-item>
        <img src="../imgs/0.png" alt="">
      </el-carousel-item>
      <el-carousel-item>
        <img src="../imgs/1.png" alt="">
      </el-carousel-item>
      <el-carousel-item>
        <img src="../imgs/2.png" alt="">
      </el-carousel-item>
      <el-carousel-item>
        <img src="../imgs/3.png" alt="">
      </el-carousel-item>
      <el-carousel-item>
        <img src="../imgs/4.png" alt="">
      </el-carousel-item>
    </el-carousel>
  </div>
</template>
<script>
export default {};
</script>
<style lang="less" scoped>
.el-carousel__item {
  img {
    width: 100%;
    height: 400px;
  }
}
.home-carousel {
  margin-top: 20px;
}
</style>
