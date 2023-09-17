
<template>
  <div>
    <Search @results="updateEmails" />
    <EmailList :emails="emails" />
  </div>
</template>

<script>
import Search from './components/Search.vue';
import EmailList from './components/EmailList.vue';

export default {
  components: {
    Search,
    EmailList,
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