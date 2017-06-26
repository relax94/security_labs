using System;
using System.Collections.Generic;
using System.Globalization;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Text.RegularExpressions;
using System.Threading.Tasks;

namespace interactive_viz
{
    class Program
    {
      static string text = "F96DE8C227A259C87EE1DA2AED57C93FE5DA36ED4EC87EF2C63AAE5B9A7EFFD673BE4ACF7BE8923CAB1ECE7AF2DA3DA44FCF7AE29235A24C963FF0DF3CA3599A70E5DA36BF1ECE77F8DC34BE129A6CF4D126BF5B9A7CFEDF3EB850D37CF0C63AA2509A76FF9227A55B9A6FE3D720A850D97AB1DD35ED5FCE6BF0D138A84CC931B1F121B44ECE70F6C032BD56C33FF9D320ED5CDF7AFF9226BE5BDE3FF7DD21ED56CF71F5C036A94D963FF8D473A351CE3FE5DA3CB84DDB71F5C17FED51DC3FE8D732BF4D963FF3C727ED4AC87EF5DB27A451D47EFD9230BF47CA6BFEC12ABE4ADF72E29224A84CDF3FF5D720A459D47AF59232A35A9A7AE7D33FB85FCE7AF5923AA31EDB3FF7D33ABF52C33FF0D673A551D93FFCD33DA35BC831B1F43CBF1EDF67F0DF23A15B963FE5DA36ED68D378F4DC36BF5B9A7AFFD121B44ECE76FEDC73BE5DD27AFCD773BA5FC93FE5DA3CB859D26BB1C63CED5CDF3FE2D730B84CDF3FF7DD21ED5ADF7CF0D636BE1EDB79E5D721ED57CE3FE6D320ED57D469F4DC27A85A963FF3C727ED49DF3FFFDD24ED55D470E69E73AC50DE3FE5DA3ABE1EDF67F4C030A44DDF3FF5D73EA250C96BE3D327A84D963FE5DA32B91ED36BB1D132A31ED87AB1D021A255DF71B1C436BF479A7AF0C13AA14794";
       //private static string text = "436f6d707574657220736369656e636520697320746865207374756479206f6620746865207468656f72792c206578706572696d656e746174696f6e2c20616e6420656e67696e656572696e67207468617420666f726d2074686520626173697320666f72207468652064657369676e20616e6420757365206f6620636f6d7075746572732e2049742069732074686520736369656e746966696320616e642070726163746963616c20617070726f61636820746f20636f6d7075746174696f6e20616e6420697473206170706c69636174696f6e7320616e64207468652073797374656d61746963207374756479206f662074686520666561736962696c6974792c207374727563747572652c2065787072657373696f6e2c20616e64206d656368616e697a6174696f6e206f6620746865206d6574686f646963616c2070726f6365647572657320286f7220616c676f726974686d7329207468617420756e6465726c696520746865206163717569736974696f6e2c20726570726573656e746174696f6e2c2070726f63657373696e672c2073746f726167652c20636f6d6d756e69636174696f6e206f662c20616e642061636365737320746f20696e666f726d6174696f6e2e20416e20616c7465726e6174652c206d6f72652073756363696e637420646566696e6974696f6e206f6620636f6d707574657220736369656e636520697320746865207374756479206f66206175746f6d6174696e6720616c676f726974686d69632070726f6365737365732074686174207363616c652e204120636f6d707574657220736369656e74697374207370656369616c697a657320696e20746865207468656f7279206f6620636f6d7075746174696f6e20616e64207468652064657369676e206f6620636f6d7075746174696f6e616c2073797374656d73";

        static HashSet<int> filterBytes(List<string> bytes, int start, int offset)
        {
            HashSet<int> filteredBytesList = new HashSet<int>();
            for (int w = start; w < bytes.Count; w += offset)
                filteredBytesList.Add(int.Parse(bytes[w], NumberStyles.HexNumber));
            return filteredBytesList;
        }

        static string openByKeyValue(List<string> bytes, int keyChangePosition, int keyLength, int keyCandidateValue)
        {
            for (int i = 0; i < bytes.Count - 1; i += keyLength)
            {
                var x = int.Parse(bytes[i + keyChangePosition], NumberStyles.HexNumber);
                var y = keyCandidateValue;

                var newByte = x ^ y;
                bytes[i + keyChangePosition] = "" + (char)newByte;
            }

            return string.Join(string.Empty, bytes);
        }

        //static bool analyze(HashSet<char> symbolicQuery)
        //{
            
        //}

        static int[] getKeysVariants(List<string> bytes, List<HashSet<char>> candidatesList, int keyLength)
        {

            List<List<char>> symbolic = new List<List<char>>();

            int[] keysCandidates = new int[keyLength];

            for (int l = 0; l < keyLength; l++)
            {
                for (int k = 0; k < candidatesList.ElementAt(l).Count; k++)
                {
                    var set = filterBytes(bytes, l, keyLength);
                    symbolic.Add(new List<char>());
                    for (int i = 0; i < set.Count; i++)
                    {
                        var a = set.ElementAt(i);
                        var b = (int) candidatesList.ElementAt(l).ElementAt(k);
                        symbolic.ElementAt(k).Add((char) (a ^ b));
                    }

                }
                keysCandidates[l] = candidatesList[l].Count > 0 ? candidatesList[l].ElementAt(0) : 0;  // MAKE ALGORITHM STRONGER
            }

            return keysCandidates;
        }

        static List<HashSet<char>> makeCandidatesList(List<string> bytes, int keyLength)
        {
            List<HashSet<char>> candidatesList = new List<HashSet<char>>();
            for (int k = 0; k < keyLength; k++)
            {
                candidatesList.Add(new HashSet<char>());
                var set = filterBytes(bytes, k, keyLength);
                for (int i = 0; i < 256; i++)
                {
                    if (set.All(x =>
                    {
                        var response = x ^ i;
                        if ((response == 32) || (response >= 35 && response < 60) || (response >= 65 && response < 90) || (response >= 97 && response < 122))
                            return true;
                        return false;

                    }))
                    {
                        candidatesList.ElementAt(k).Add((char)i);
                    }
                }

            }
            return candidatesList;
        }

        static Dictionary<int, double> probateKeyLength()
        {
            Dictionary<int, double> stats = new Dictionary<int, double>();

            for (int i = 2; i < 20; i++)
            {
                var query = (from Match m in Regex.Matches(text, @".{1," + (i * 2) + "}")
                             select m.Value).ToList();
                var group = query.Select(x => x.Substring(0, 2)).GroupBy(k => k);
                var L = query.Count();
                var S = group.Sum(g =>
                {
                    var c = g.Count();
                    return Math.Pow((double)c / L, 2);
                });

                stats.Add(i, S);
            }
            return stats;
        }


        static void Main(string[] args)
        {
            var bytes = (from Match m in Regex.Matches(text, @".{1,2}")
                         select m.Value).ToList();


            var order =  probateKeyLength().OrderByDescending(o => o.Value);

            foreach (var orderItem in order)
            {
              Console.WriteLine($"Key length = {orderItem.Key}    S = {orderItem.Value}");
            }
            Console.WriteLine();
            Console.Write("Probate keyLength = ");

            int keyLength = int.Parse(Console.ReadLine());


            var candidatesList = makeCandidatesList(bytes, keyLength);

            var keyCandidate = getKeysVariants(bytes, candidatesList, keyLength);

            Console.ForegroundColor = ConsoleColor.Green;
            Console.WriteLine($"\n\n KEY BYTE DEC SNAPSHOT : {String.Join(" ", keyCandidate)}\nPress any key for step by step flow\n");
            Console.ResetColor();

            //var keyCandidate = new int[7] { 186, 31, 145, 178, 83, 205, 62 };
         //   var keyCandidate = new int[7] { 105, 104, 100, 109,105,116,114 };

            for (int i = 0; i < keyCandidate.Length; i++)
            {
                Console.WriteLine(openByKeyValue(bytes, i, keyLength, keyCandidate[i]));
                Console.WriteLine();
                Console.WriteLine();
                Console.WriteLine("Press Space To Start Encrypt Cycle");
                Console.ReadKey();
                Console.WriteLine();
                Console.WriteLine();
            }



        }
    }
}
