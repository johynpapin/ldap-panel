<template>
  <v-row no-gutters justify="center">
    <v-col cols="12" md="10" lg="8" xl="6">
      <h1 class="text-h4 mb-8">Changement de mot de passe</h1>

      <ValidationObserver ref="form" v-slot:default="{ handleSubmit, invalid }">
        <form @submit.prevent="handleSubmit(updatePassword)">
          <ValidationProvider name="Nouveau mot de passe" rules="required" vid="password" v-slot:default="{ errors }">
            <v-text-field
              v-model="password"
              :error-messages="errors"
              type="password"
              label="Nouveau mot de passe"
              outlined
            ></v-text-field>
          </ValidationProvider>

          <ValidationProvider name="Vérification du mot de passe" rules="required|confirmed:password" v-slot:default="{ errors }">
            <v-text-field
              v-model="passwordVerification"
              :error-messages="errors"
              type="password"
              label="Vérification du mot de passe"
              outlined
            ></v-text-field>
          </ValidationProvider>

          <v-row justify="end">
            <v-col cols="12" sm="auto" class="flex-sm-shrink-1">
              <v-btn
                :disabled="updatingPassword || invalid"
                :loading="updatingPassword"
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
</template>

<script>
import { ValidationObserver, ValidationProvider } from 'vee-validate'

export default {
  name: 'UpdatePasswordView',

  metaInfo: {
    title: 'Changement de mot de passe'
  },

  data: () => ({
    updatingPassword: false,

    password: '',
    passwordVerification: ''
  }),

  methods: {
    async updatePassword () {
      this.updatingPassword = true

      try {
        await this.$store.dispatch('data/updatePassword', {
          password: this.password
        })

        this.password = ''
        this.passwordVerification = ''

        this.$nextTick(() => {
          this.$refs.form.reset()
        })
      } finally {
        this.updatingPassword = false
      }
    }
  },

  components: {
    ValidationObserver,
    ValidationProvider
  }
}
</script>
