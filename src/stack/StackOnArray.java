package stack;


class ArrayStack<T> {
    private T[] data;
    private int top;
    private int capacity;

    private static int d4tCap = 32;

    public ArrayStack(final int capa) {
        data = (T[]) new Object[capa];
        capacity = capa;
        top = -1;
    }

    public ArrayStack() {
        data = (T[]) new Object[d4tCap];
        capacity = d4tCap;
        top = -1;
    }

    public void Push(final T item) {
        if (top < 0) {
            top = 0;
        }

        if (top == capacity - 1) {
            resize();
        }

        top++;
        data[top] = item;

    }

    public T Pop() {
        if (top <= 0) {
            return null;
        }

        final T v = data[top];
        top--;
        return v;
    }

    public void Flush() {
        top = -1;

    }

    public T Top() {
        return (top > -1) ? data[top] : null;
    }

    public void resize() {
        T[] newdata = (T[]) new Object[capacity*2];

        for (int i = 0 ; i < capacity; i ++) {
            newdata[i] = data[i];
        }

        capacity = capacity*2;
        data = newdata;
    }
}