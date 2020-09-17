import { configure, extend, setInteractionMode, localize } from 'vee-validate'
import fr from 'vee-validate/dist/locale/fr.json'
import { required, confirmed } from 'vee-validate/dist/rules'

localize('fr', fr)

setInteractionMode('aggressive')

configure({})

extend('required', required)
extend('confirmed', confirmed)
