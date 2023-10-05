<template>
  <div>
    <Search @updateTerm="term = $event" @searched="updateEmailContents" />
    <EmailList :term="term" :emailContents="emailContents" :totalHits="totalHits" />
    <Pagination :totalHits="totalHits" />


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
      totalHits: 0  
    };
  },
  methods: {
    updateEmailContents({ emailContents, totalHits }) {
    this.emailContents = emailContents;
    this.totalHits = totalHits;
  }
  },

  methods: {
    updateEmailContents({ emailContents, totalHits }) {
    this.emailContents = emailContents;
    this.totalHits = totalHits;
  }
  }
};
</script>
