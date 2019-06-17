import Configuration from '@/config'

export default {
    computed: {
        config () {
            return Configuration
        }
    },
    methods: {
        formatString (str) {
            for (let i = 1; i < arguments.length; i++) {
              const replacement = '{' + (i - 1) + '}'
              str = str.replace(replacement, arguments[i])
            }
            return str
        }
    }
}
