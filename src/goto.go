func foo int(bar int) {
    return_value := 0

    if ret1, err := do_something(bar); err {
        goto error_1 // HL
    }

    if ret2, err := init_stuff(ret1); err {
        goto error_2 // HL
    }

    if ret3, err := prepare_stuff(ret2); err {
        goto error_3 // HL
    }

    return_value = do_the_thing(ret3)
error_3: // HL
    cleanup_3()
error_2: // HL
    cleanup_2()
error_1: // HL
    cleanup_1()
    return return_value
}
