import {
  Stack,
  Container,
  Text,
  Heading,
  Accordion,
  AccordionButton,
  AccordionItem,
  AccordionIcon,
  AccordionPanel,
  Box,
  useColorModeValue,
} from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import axios from 'axios';
import { ENDPOINT_API_GET_FAQ } from '../../../api/api';

export default function Faq() {
  const [faq, setFaq] = useState([]);
  useEffect(() => {
    const fetchData = async () => {
      const result = await axios.get(ENDPOINT_API_GET_FAQ);
      setFaq(result.data);
    };
    fetchData();
  }, []);

  const AccordionContainer = (props) => (
    <AccordionItem border="0">
      <AccordionButton px={0} _hover={{ bg: 'none' }}>
        <Box flex="1" textAlign="left">
          {props.tanya}
        </Box>
        <AccordionIcon ml={5} />
      </AccordionButton>
      <AccordionPanel pb={4} px={0}>
        <Text color={'gray.500'} align="justify">
          {props.jawab}
        </Text>
      </AccordionPanel>
    </AccordionItem>
  );

  return (
    <Stack
      as={Container}
      mt={20}
      mb={6}
      maxW={'7xl'}
      align="center"
      justifyContent="center"
      minH={{ base: 'auto', lg: '70vh' }}
      data-aos="fade-up"
    >
      <Stack
        p={{ base: 5, lg: 10 }}
        bg={useColorModeValue('gray.100', 'gray.900')}
        rounded="xl"
        direction={{ base: 'column', md: 'row' }}
        spacing={{ base: 4, md: 10 }}
      >
        <Stack flex={1}>
          <Stack>
            <Text
              color={useColorModeValue('accentLight.500', 'accentDark.500')}
            >
              Support
            </Text>
            <Heading as="h2">FAQ</Heading>
            {faq.map((faqs) => (
              <Text fontSize={'md'} color={'gray.500'} align="justify">
                {faqs.judul}
              </Text>
            ))}
          </Stack>
        </Stack>
        <Stack flex={1}>
          <Accordion allowToggle w="full">
            {faq.map((item, index) => (
              <AccordionContainer key={index} {...item} />
            ))}
          </Accordion>
        </Stack>
      </Stack>
    </Stack>
  );
}
