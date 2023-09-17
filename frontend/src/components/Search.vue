<template>
    <div class="flex items-center border rounded-lg shadow-md">
      <input v-model="term" placeholder="Buscar emails..." class="flex-grow py-2 px-4 outline-none"/>
      <button @click="search" class="px-4 text-gray-500 focus:outline-none">
        <i class="fas fa-search"></i>
      </button>
      <span class="px-4 text-gray-500">
        <i class="fas fa-filter"></i>
      </span>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        term: '',
        results: []
      };
    },
    methods: {
      async search() {
        if (!this.term) {
          this.results = [];
          return;
        }
  
        try {
          const response = await fetch(`http://localhost:8080/api/search?term=${this.term}`);
  const data = await response.json();
  this.$emit('results', data.hits.hits);
  
        } catch (error) {
          console.error("Error fetching search results:", error);
        }
      }
    }
  };
  </script>