<template>
    <div class=" m-auto">
      <div :id="id" :option="option"></div>
    </div>
</template>

<script>
  import HighCharts from 'highcharts'

  export default {
    name: 'PieChart',
    components: {
      HighCharts
    },
    props: {
      id: null,
      option: {
        type: Object
      }
    },
    mounted() {
      HighCharts.chart(this.id, this.option)
    }
  }
</script>

<style scoped>

</style>
