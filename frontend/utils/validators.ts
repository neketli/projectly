type RuleTypes = 'string' | 'number' | 'email'

export const validators = {
    required: {
        required: true,
        message: 'Поле является обязательным',
        trigger: ['blur', 'change'],
    },
    requiredNumber: {
        type: 'number' as RuleTypes,
        required: true,
        message: 'Поле является обязательным',
        trigger: ['blur', 'change'],
    },
    len: (min = 2, max = 128) => ({
        type: 'string' as RuleTypes,
        min,
        max,
        message: `Длинна поля должна быть не меньше ${min} и не больше ${max}`,
        trigger: 'blur',
    }),
    range: (min?: number, max?: number) => ({
        type: 'number' as RuleTypes,
        min,
        max,
        message: `Поле должно быть числом в диапазоне от ${min} и до ${max}`,
        trigger: 'blur',
    }),
    email: {
        type: 'email' as RuleTypes,
        message: 'Введите корректный email',
        trigger: 'blur',
    },
    cyrillic: {
        type: 'string' as RuleTypes,
        validator: (_: unknown, value: string, callback: (error?: Error) => unknown) => {
            if (!/^[а-яА-ЯёЁ\s]+$/.test(value)) {
                callback(new Error('Поле должно состоять из символов кириллицы'))
            }
            else {
                callback()
            }
        },
        trigger: 'blur',
    },
    phone: {
        type: 'string' as RuleTypes,
        validator: (_: unknown, value: string, callback: (error?: Error) => unknown) => {
            if (!/^((\+7|7|8)+([0-9]){10})$/.test(value)) {
                callback(
                    new Error(
                        'Поле должно содержать российский номер телефона, без пробелов и прочих символов',
                    ),
                )
            }
            else {
                callback()
            }
        },
        trigger: 'blur',
    },
}
