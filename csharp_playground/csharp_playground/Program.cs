using System;

namespace csharp_playground
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
            var t = new T();
            Console.WriteLine(t.s);
            c(t);
            Console.WriteLine(t.s);
        }
        static void c(T a1)
        {
            var a2 = new T();
            a2.s = "abc";
            a1 = a2;
        }
    }

    class T
    {
        public string s = "123";
    }
}
