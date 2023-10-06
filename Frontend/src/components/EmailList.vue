

<template>
    <div >

    <!-- Tabla -->
    <table class="table-auto w-full text-sm text-left text-gray-500 dark:text-gray-400 ">
      <!-- Encabezados de la tabla -->
      <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
        <tr>
          <th class="px-6 py-3">Sent by</th>
          <th class="px-6 py-3">Subject</th>
          <th class="px-6 py-3">Preview</th>
          <th class="px-6 py-3">Date</th>
        </tr>
      </thead>
      <tbody >
        <!-- Lista de correos -->
        <tr class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-sky-900 rounded-xl cursor-pointer" v-for="emailContent in emailContents" :key="emailContent._id" @click="selectEmail(emailContent)">
          <td class="px-6 py-4 border font-medium text-gray-900 whitespace-nowrap dark:text-white">
            {{ extractField(emailContent.content, 'From') }}
          </td>
          <td class="px-6 py-4 border" v-html="highlightKeyword(extractField(emailContent.content, 'Subject'), term)">
          </td>
          <td class="px-6 py-4 border" v-html="highlightKeyword(extractContentPreview(emailContent.content), term)">
          </td>
          <td class="px-6 py-4 border">
            {{ formatDate(extractField(emailContent.content, 'Date')) }}
          </td>
        </tr>
      </tbody>
      <EmailDetail v-if="selectedEmail" :email="selectedEmail" :term="term" />
    </table>
  </div>

</template>

<script>
import EmailDetail from './EmailDetail.vue';

export default {

  components: {
    EmailDetail,
    
  },
  props: {
    emailContents: {
      type: Array,
      required: true
    },
    term: {
    type: String,
    default: ''
  },
  totalHits: {  
      type: Number,
      default: 0
    },
  },
  
  data() {
    return {
      selectedEmail: null
    };
  },
  methods: {
    extractField(content, field) {
      if (!content) return '';
      const regex = new RegExp(`${field}: (.+)\\r\\n`);
      const match = content.match(regex);
      return match ? match[1].trim() : '';
    },

    extractContentPreview(content) {
      const body = content.split('\r\n\r\n')[1];
      return body ? body.split(' ').slice(0, 5).join(' ') + '...' : '';
    },
    formatDate(dateString) {
      const date = new Date(dateString);
      return `${date.getDate().toString().padStart(2, '0')}/${(date.getMonth() + 1).toString().padStart(2, '0')}/${date.getFullYear()}`;
    },
    selectEmail(email) {
      this.selectedEmail = email;
    },
    highlightKeyword(text, keyword) {
      if (!keyword) return text;
      const regex = new RegExp(`(${keyword})`, 'gi');
      return text.replace(regex, '<span class="highlight">$1</span>');
    }

  }
};
</script>

<style scoped>
.highlight {
  background-color: yellow;
  padding: 2px;
}
</style>



