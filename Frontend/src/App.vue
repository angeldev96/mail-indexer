<template>
  <div>
    <Search @updateTerm="term = $event" @searched="updateEmailContents" :currentPage="currentPage" />
    <EmailList :term="term" :emailContents="emailContents" :totalHits="totalHits" />
    <Pagination :totalHits="totalHits" @pageChanged="handlePageChange" />


  </div>
</template>

<script>
import Search from './components/Search.vue';
import EmailList from './components/EmailList.vue'
import Pagination from './components/Pagination.vue';


export default {
  components: {
    Search,
    EmailList,
    Pagination

  },
  data() {
  return {
    emailContents: [],
    term: '',
    totalHits: 0,
    currentPage: 1  // Añade esta línea
  };
},

  methods: {
    updateEmailContents({ emailContents, totalHits }) {
      this.emailContents = emailContents;
      this.totalHits = totalHits;
    },
    handlePageChange(page) {
      this.currentPage = page;
      // Aquí puedes hacer la solicitud de búsqueda nuevamente con la nueva página.
      // Puedes emitir un evento desde aquí al componente `Search` para que realice la búsqueda
      // o mover la lógica de búsqueda aquí.
   },
  }
};
</script>
