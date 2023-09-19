
<template>
  <div >
    <Search @results="updateEmails" />
    <EmailList :emails="emails" />
  </div>
</template>

<script>
import Search from './components/Search.vue';

export default {
  components: {
    Search,
  },
  data() {
    return {
      emails: []
    };
  },
  methods: {
    updateEmails(results) {
      this.emails = results;
    }
  }
};
</script>