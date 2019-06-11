using System;
using System.Numerics;
using System.Security.Cryptography;

namespace csharp_playground
{
    class Program
    {
        static void Main(string[] args)
        {

        }
    }

    class RSA
    {
        private BigInteger p { get; set; }
        private BigInteger q { get; set; }
        private BigInteger n { get; set; }
        private BigInteger phi { get; set; }
        private BigInteger e { get; set; }
        private BigInteger d { get; set; }
        private RandomNumberGenerator rng { get; set; }
        private int bitLength { get; set; }
        private int blockSize { get; set; }

        RSA()
        {
            rng = new RNGCryptoServiceProvider();
            p.
        }
    }
}
