<template>
  <div>
    <!-- Componente de búsqueda -->
    <Search @searched="updateEmailContents" />

    <!-- Componente para mostrar los correos -->
    <EmailList :emailContents="emailContents" />
  </div>
</template>

<script>
import Search from './components/Search.vue';
import EmailList from './components/EmailList.vue'

export default {
  components: {
    Search,
    EmailList
  },
  data() {
    return {
      emailContents: []
    };
  },
  methods: {
    updateEmailContents(newEmailContents) {
      this.emailContents = newEmailContents;
    }
  }
};
</script>
