<template>
  <div class="about">
    <el-row class="about-item" :gutter="10">
      <el-col :xs="2" :sm="6" :md="2" :lg="2" :xl="1">地址:</el-col>
      <el-col :xs="10" :sm="6" :md="10" :lg="10" :xl="11">山东省潍坊市青州市益都明祖山路669号</el-col>
    </el-row>
    <el-row class="about-item" :gutter="10">
      <el-col :xs="2" :sm="6" :md="2" :lg="2" :xl="1">电话:</el-col>
      <el-col :xs="10" :sm="6" :md="10" :lg="10" :xl="11">13606470036</el-col>
    </el-row>
    <el-row class="about-item" :gutter="10">
      <el-col :xs="2" :sm="6" :md="2" :lg="2" :xl="1">邮箱:</el-col>
      <el-col :xs="10" :sm="6" :md="10" :lg="10" :xl="11">61137700@qq.com</el-col>
    </el-row>
  </div>
</template>
<style lang="less" scoped>
  .about{
    padding: 40px 0 40px;
    .about-item{
      padding: 10px 0;
    }
  }
</style>
