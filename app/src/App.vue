<template>
  <v-app>
    <RouterView></RouterView>

    <AlertSnackbar></AlertSnackbar>
  </v-app>
</template>

<script>
import AlertSnackbar from '@/components/AlertSnackbar'

export default {
  name: 'App',

  metaInfo: {
    title: 'Accueil',
    titleTemplate: 'Panel LDAP - %s'
  },

  components: {
    AlertSnackbar
  }
}
</script>
