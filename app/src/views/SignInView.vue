<template>
  <v-main>
    <v-container>
      <v-row justify="center">
        <v-col cols="12" sm="10" lg="6">
          <h1 class="text-h4 mb-8">Connexion</h1>

          <ValidationObserver v-slot:default="{ handleSubmit, invalid }">
            <form @submit.prevent="handleSubmit(signIn)">
              <ValidationProvider name="Identifiant" rules="required" v-slot:default="{ errors }">
                <v-text-field
                  v-model="username"
                  :error-messages="errors"
                  label="Identifiant"
                  outlined
                ></v-text-field>
              </ValidationProvider>

              <ValidationProvider name="Mot de passe" rules="required" v-slot:default="{ errors }">
                <v-text-field
                  v-model="password"
                  :error-messages="errors"
                  type="password"
                  label="Mot de passe"
                  outlined
                ></v-text-field>
              </ValidationProvider>

              <v-row justify="end">
                <v-col cols="12" sm="auto" class="flex-sm-shrink-1">
                  <v-btn
                    :disabled="signingIn || invalid"
                    :loading="signingIn"
                    type="submit"
                    color="primary"
                    large
                    block
                  >Envoyer</v-btn>
                </v-col>
              </v-row>
            </form>
          </ValidationObserver>
        </v-col>
      </v-row>
    </v-container>
  </v-main>
</template>

<script>
import { ValidationObserver, ValidationProvider } from 'vee-validate'

export default {
  name: 'SignInView',

  metaInfo: {
    title: 'Connexion'
  },

  data: () => ({
    signingIn: false,

    username: '',
    password: ''
  }),

  methods: {
    async signIn () {
      this.signingIn = true

      try {
        await this.$store.dispatch('data/signIn', {
          username: this.username,
          password: this.password
        })
      } finally {
        this.signingIn = false
      }
    }
  },

  components: {
    ValidationObserver,
    ValidationProvider
  }
}
</script>
