<script setup>
import { computed } from 'vue';

const props = defineProps({
  data: Object
});

const formattedSeriesData = computed(() => {
  return [{
    data: Object.entries(props.data)
      .filter(([name, value]) => value !== 1) 
      .map(([name, value]) => ({
        x: name, // назва елемента
        y: value  // значення елемента
      }))
  }];
});



const options = {
  chart: {
    type: 'treemap',
    toolbar: {
      show: true
    }
  },
  title: {
    text: 'TreeMap Chart '
  },
  colors: ['#3B93A5', '#F7B844', '#ADD8C7', '#EC3C65', '#CDD7B6'],
  plotOptions: {
    treemap: {
      distributed: true,
      enableShades: true,
      shadeIntensity: 0.5,
      reverseNegativeShade: true
    }
  }
}
</script>

<template>
<div>
  <apexchart width="1200" :options="options" :series="formattedSeriesData"></apexchart>
</div>
</template>
