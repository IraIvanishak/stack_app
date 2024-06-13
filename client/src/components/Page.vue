<template>
    <div>
      <div class="toolbar">
        <label for="category">Select your specialization:</label>
        <select v-model="selectedCategory" @change="fetchData">
          <option value="Golang">Golang</option>
          <option value="PHP">PHP</option>
          <option value="Python">Python</option>
        </select>
      </div>
      <tree-map-chart :data="ata"></tree-map-chart>
    </div>
  </template>
  
  <script>
  import TreeMapChart from './TreeMapChart.vue';
  
  export default {
    components: {
      TreeMapChart
    },
    data() {
      return {
        ata: {}, // Початковий стан для даних
        selectedCategory: 'Golang', // Значення за замовчуванням для селекту категорії
        someBoolean: false // Булева змінна для чекбоксу
      }
    },
    mounted() {
      this.fetchData();
    },
    methods: {
      fetchData() {
        fetch(`http://localhost:8880/analytic?category=${this.selectedCategory}`)
          .then(response => {
            if (!response.ok) {
              throw new Error(`HTTP error! status: ${response.status}`);
            }
            return response.json();
          })
          .then(data => {
                        this.ata = data; // Записуємо отримані дані у локальний стан
          })
          .catch(error => {
            console.error('There was an error fetching the data:', error);
          });
      }
    }
  }
  </script>
  

<style>
label {
    color: white;
  font-size: 16px;
  margin-right: 10px;
}
.toolbar {
  padding: 10px;
  background-color: #f4f4f46b; /* Світлий фон для панелі */
  border-radius: 5px;
  margin-bottom: 20px;
}

.toolbar select, .toolbar input[type="checkbox"] {
    
    font-size: 16px;
    padding: 8px;
  margin-right: 10px;
}
</style>

