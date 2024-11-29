type RuleTypes = 'string' | 'number' | 'email'

export const useValidator = () => {
    const { t } = useI18n()

    return {
        required: {
            required: true,
            message: t('common.validation.required'),
            trigger: ['blur', 'change'],
        },
        requiredNumber: {
            type: 'number' as RuleTypes,
            required: true,
            message: t('common.validation.required'),
            trigger: ['blur', 'change'],
        },
        len: (min = 2, max = 128) => ({
            type: 'string' as RuleTypes,
            min,
            max,
            message: t('common.validation.length', { min, max }),
            trigger: 'blur',
        }),
        range: (min?: number, max?: number) => ({
            type: 'number' as RuleTypes,
            min,
            max,
            message: t('common.validation.range', { min, max }),
            trigger: 'blur',
        }),
        email: {
            type: 'email' as RuleTypes,
            message: t('common.validation.email'),
            trigger: 'blur',
        },
    }
}
